package consensus

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"net/netip"
	"strconv"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	abcitypes "github.com/cometbft/cometbft/abci/types"
	v1 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	cmtproxy "github.com/cometbft/cometbft/proxy"
	cmttypes "github.com/cometbft/cometbft/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/meterio/supernova/block"
	"github.com/meterio/supernova/chain"
	"github.com/meterio/supernova/types"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls"
)

var (
	ErrInvalidBlock = errors.New("invalid block")
)

type Executor struct {
	proxyApp cmtproxy.AppConnConsensus
	chain    *chain.Chain
	logger   *slog.Logger
	eventBus cmttypes.BlockEventPublisher
}

func NewExecutor(proxyApp cmtproxy.AppConnConsensus, c *chain.Chain) *Executor {
	return &Executor{proxyApp: proxyApp, chain: c, logger: slog.With("pkg", "exec")}
}

func (e *Executor) InitChain(req *abcitypes.InitChainRequest) (*abcitypes.InitChainResponse, error) {
	return e.proxyApp.InitChain(context.TODO(), req)
}

func (e *Executor) PrepareProposal(parent *block.DraftBlock, proposerIndex int) (*abcitypes.PrepareProposalResponse, error) {
	maxBytes := int64(cmttypes.MaxBlockSizeBytes)

	evSize := int64(0)
	vset := e.chain.GetValidatorsByHash(parent.ProposedBlock.NextValidatorsHash())
	maxDataBytes := cmttypes.MaxDataBytes(maxBytes, evSize, vset.Size())
	return e.proxyApp.PrepareProposal(context.TODO(), &v1.PrepareProposalRequest{
		MaxTxBytes:         maxDataBytes,
		Height:             int64(parent.Height) + 1,
		Time:               time.Now(),
		Misbehavior:        make([]v1.Misbehavior, 0), // FIXME: track the misbehavior and preppare the evidence
		NextValidatorsHash: parent.ProposedBlock.NextValidatorsHash(),
		ProposerAddress:    vset.GetByIndex(proposerIndex).Address.Bytes(),
	})
}

func (e *Executor) ProcessProposal(blk *block.Block) (bool, error) {
	vset := e.chain.GetValidatorsByHash(blk.ValidatorsHash())
	parent, err := e.chain.GetBlock(blk.ParentID())
	if err != nil {
		parentDraft := e.chain.GetDraft(blk.ParentID())
		parent = parentDraft.ProposedBlock
	}
	resp, err := e.proxyApp.ProcessProposal(context.TODO(), &v1.ProcessProposalRequest{
		Hash:               blk.ID().Bytes(),
		Height:             int64(blk.Number()),
		Time:               time.Unix(int64(blk.Timestamp()), 0),
		Txs:                blk.Txs.Convert(),
		ProposedLastCommit: e.chain.BuildLastCommitInfo(parent),
		Misbehavior:        make([]v1.Misbehavior, 0), // FIXME: track the misbehavior and preppare the evidence
		ProposerAddress:    vset.GetByIndex(int(blk.ProposerIndex())).Address.Bytes(),
		NextValidatorsHash: blk.NextValidatorsHash(),
	})

	if err != nil {
		return false, err
	}
	if resp.IsStatusUnknown() {
		panic("ProcessProposal responded with status " + resp.Status.String())
	}
	return resp.IsAccepted(), nil
}

func (e *Executor) ExtendVote(req *abcitypes.ExtendVoteRequest) (*abcitypes.ExtendVoteResponse, error) {
	return e.proxyApp.ExtendVote(context.TODO(), req)
}

func (e *Executor) VerifyVoteExtension(req *abcitypes.VerifyVoteExtensionRequest) (*abcitypes.VerifyVoteExtensionResponse, error) {
	return e.proxyApp.VerifyVoteExtension(context.TODO(), req)
}

func (e *Executor) FinalizeBlock(req *abcitypes.FinalizeBlockRequest) (*abcitypes.FinalizeBlockResponse, error) {
	return e.proxyApp.FinalizeBlock(context.TODO(), req)
}

func (e *Executor) Commit() (*abcitypes.CommitResponse, error) {
	return e.proxyApp.Commit(context.TODO())
}

func validateBlock(b *block.Block) error {
	// FIXME: imple this
	return nil
}

// ApplyBlock validates the block against the state, executes it against the app,
// fires the relevant events, commits the app, and saves the new state and responses.
// It returns the new state.
// It's the only function that needs to be called
// from outside this package to process and commit an entire block.
// It takes a blockID to avoid recomputing the parts hash.
func (e *Executor) ApplyBlock(block *block.Block, syncingToHeight int64) ([]byte, *types.ValidatorSet, error) {
	if err := validateBlock(block); err != nil {
		return make([]byte, 0), nil, ErrInvalidBlock
	}

	return e.applyBlock(block, syncingToHeight)
}

func (e *Executor) applyBlock(blk *block.Block, syncingToHeight int64) (appHash []byte, nxtVSet *types.ValidatorSet, err error) {
	vset := e.chain.GetValidatorsByHash(blk.ValidatorsHash())
	parent, err := e.chain.GetBlock(blk.ParentID())
	if err != nil {
		parentDraft := e.chain.GetDraft(blk.ParentID())
		parent = parentDraft.ProposedBlock
	}
	abciResponse, err := e.proxyApp.FinalizeBlock(context.TODO(), &abci.FinalizeBlockRequest{
		Hash:               blk.ID().Bytes(),
		NextValidatorsHash: blk.Header().NextValidatorsHash,
		ProposerAddress:    vset.GetByIndex(int(blk.ProposerIndex())).Address.Bytes(),
		Height:             int64(blk.Number()),
		Time:               time.Unix(int64(blk.Timestamp()), 0),
		DecidedLastCommit:  e.chain.BuildLastCommitInfo(parent),
		Misbehavior:        make([]v1.Misbehavior, 0), // FIXME: track the misbehavior and preppare the evidence
		Txs:                blk.Transactions().Convert(),
		SyncingToHeight:    syncingToHeight,
	})
	appHash = abciResponse.AppHash
	e.logger.Info(
		"Finalized block",
		"height", blk.Number(),
		"num_txs_res", len(abciResponse.TxResults),
		"num_val_updates", len(abciResponse.ValidatorUpdates),
		"block_app_hash", fmt.Sprintf("%X", abciResponse.AppHash),
		"syncing_to_height", syncingToHeight,
	)

	// Assert that the application correctly returned tx results for each of the transactions provided in the block
	if len(blk.Txs) != len(abciResponse.TxResults) {
		err = fmt.Errorf("expected tx results length to match size of transactions in block. Expected %d, got %d", len(blk.Txs), len(abciResponse.TxResults))
		return
	}

	// calculate the next committee
	if len(abciResponse.ValidatorUpdates) > 0 {
		curVSet := e.chain.GetValidatorsByHash(blk.ValidatorsHash())
		nxtVSet = calcNewValidatorSet(curVSet, abciResponse.ValidatorUpdates, abciResponse.Events)
	} else {
		nxtVSet = nil
	}

	return
}

func calcNewValidatorSet(vset *types.ValidatorSet, updates abcitypes.ValidatorUpdates, events []abcitypes.Event) (nxtVSet *types.ValidatorSet) {
	if updates.Len() <= 0 {
		return
	}
	nxtVSet = vset.Copy()

	veMap := make(map[string]validatorExtra)
	for _, ev := range events {
		if ev.Type == "ValidatorExtra" {
			ve := validatorExtra{}
			for _, attr := range ev.Attributes {
				switch attr.Key {
				case "address":
					ve.Address = common.Address{}
				case "name":
					ve.Name = attr.Value
				case "pubkey":
					ve.Pubkey, _ = hex.DecodeString(attr.Value)
				case "ip":
					ve.IP = attr.Value
				case "port":
					ve.Port, _ = strconv.ParseUint(attr.Value, 10, 32)
				}
			}
			veMap[hex.EncodeToString(ve.Pubkey)] = ve
		}
	}
	for _, update := range updates {
		pubkey, err := bls.PublicKeyFromBytes(update.PubKeyBytes)
		pubkeyHex := hex.EncodeToString(update.PubKeyBytes)
		if err != nil {
			panic(err)
		}
		if update.Power == 0 {
			nxtVSet.DeleteByPubkey(pubkey)
		} else {
			v := nxtVSet.GetByPubkey(pubkey)

			if v == nil {
				v = &types.Validator{PubKey: pubkey, VotingPower: uint64(update.Power)}
			}

			if ve, existed := veMap[pubkeyHex]; existed {
				v.IP, err = netip.ParseAddr(ve.IP)
				if err != nil {
					continue
				}
				if ve.Port != 0 {
					v.Port = uint32(ve.Port)
				}
				if ve.Name != "" {
					v.Name = ve.Name
				}
				v.Address = common.Address{}
			}

			v.VotingPower = uint64(update.Power)
			nxtVSet.Upsert(v)
		}
	}
	return
}

func CalcAddedValidators(curVSet, nxtVSet *types.ValidatorSet) (added []*types.Validator) {
	if nxtVSet == nil {
		return
	}
	visited := make(map[string]bool)
	for _, v := range curVSet.Validators {
		visited[hex.EncodeToString(v.PubKey.Marshal())] = true
	}
	for _, v := range nxtVSet.Validators {
		if _, exist := visited[hex.EncodeToString(v.PubKey.Marshal())]; !exist {
			added = append(added, v)
		}
	}
	return
}

// SetEventBus - sets the event bus for publishing block related events.
// If not called, it defaults to types.NopEventBus.
func (e *Executor) SetEventBus(eventBus cmttypes.BlockEventPublisher) {
	e.eventBus = eventBus
}
