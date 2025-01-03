// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package consensus

import (
	"errors"
)

var (
	errFutureBlock              = errors.New("block in the future")
	errParentMissing            = errors.New("parent block is missing")
	errQCNodeMissing            = errors.New("qcNode is missing")
	errKnownBlock               = errors.New("block already in the chain")
	errParentHeaderMissing      = errors.New("parent header is missing")
	errDecodeParentFailed       = errors.New("decode parent failed")
	errRestartPaceMakerRequired = errors.New("restart pacemaker required")

	ErrUnrecognizedPayload = errors.New("unrecognized payload")
	ErrVersionMismatch     = errors.New("version mismatch")
	ErrMalformattedMsg     = errors.New("malformatted msg")
	ErrKnownMsg            = errors.New("known msg")
	ErrProposalRejected    = errors.New("proposal rejected")
	ErrProposalUnknown     = errors.New("proposal unknown")
	ErrForkHappened        = errors.New("fork happened")
)

type consensusError string

func (err consensusError) Error() string {
	return string(err)
}

// IsFutureBlock returns if the error indicates that the block should be
// processed later.
func IsFutureBlock(err error) bool {
	return err == errFutureBlock
}

// IsParentMissing ...
func IsParentMissing(err error) bool {
	return err == errParentMissing
}

// IsKnownBlock returns if the error means the block was already in the chain.
func IsKnownBlock(err error) bool {
	return err == errKnownBlock
}

// IsCritical returns if the error is consensus related.
func IsCritical(err error) bool {
	_, ok := err.(consensusError)
	return ok
}
