// Copyright (c) 2020 The Meter.io developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package types

import (
	"crypto/ecdsa"
	sha256 "crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls"
)

type BlsCommon struct {
	PrivKey bls.SecretKey //my private key
	PubKey  bls.PublicKey //my public key

}

func NewBlsCommon() *BlsCommon {
	secretKey, err := bls.RandKey()
	if err != nil {
		return nil
	}
	return &BlsCommon{
		PrivKey: secretKey,
		PubKey:  secretKey.PublicKey(),
	}
}

func NewBlsCommonFromParams(pubKey bls.PublicKey, privKey bls.SecretKey) *BlsCommon {
	return &BlsCommon{
		PrivKey: privKey,
		PubKey:  pubKey,
	}
}

// BLS is implemented by C, memeory need to be freed.
// Signatures also need to be freed but Not here!!!
func (cc *BlsCommon) Destroy() bool {
	return true
}

func (cc *BlsCommon) GetPublicKey() *bls.PublicKey {
	return &cc.PubKey
}

// func (cc *BlsCommon) GetPrivateKey() *bls.PrivateKey {
// 	return &cc.PrivKey
// }

// sign the part of msg
func (cc *BlsCommon) SignMessage(msg []byte) (bls.Signature, [32]byte) {
	hash := sha256.Sum256(msg)
	sig := cc.PrivKey.Sign(hash[:])
	return sig, hash
}

func (cc *BlsCommon) SignHash(hash [32]byte) []byte {
	return cc.PrivKey.Sign(hash[:]).Marshal()
}

func (cc *BlsCommon) VerifySignature(signature, msgHash, blsPK []byte) (bool, error) {
	var fixedMsgHash [32]byte
	copy(fixedMsgHash[:], msgHash[32:])
	pubkey, err := bls.PublicKeyFromBytes(blsPK)
	if err != nil {
		fmt.Println("pubkey unmarshal failed")
		return false, nil
	}
	return bls.VerifySignature(signature, [32]byte(msgHash), pubkey)
}

func (cc *BlsCommon) AggregateSign(sigs []bls.Signature) bls.Signature {
	return bls.AggregateSignatures(sigs)
}

func (cc *BlsCommon) AggregateVerify(aggrSig bls.Signature, hash [32]byte, pubkeys []bls.PublicKey) bool {
	return aggrSig.FastAggregateVerify(pubkeys, hash)
}

func (cc *BlsCommon) SplitPubKey(comboPubKey string) (*ecdsa.PublicKey, *bls.PublicKey) {
	// first part is ecdsa public, 2nd part is bls public key
	split := strings.Split(comboPubKey, ":::")
	// fmt.Println("ecdsa PubKey", split[0], "Bls PubKey", split[1])
	pubKeyBytes, err := b64.StdEncoding.DecodeString(split[0])
	if err != nil {
		panic(fmt.Sprintf("read public key of delegate failed, %v", err))
	}
	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		panic(fmt.Sprintf("read public key of delegate failed, %v", err))
	}

	blsPubBytes, err := b64.StdEncoding.DecodeString(split[1])
	if err != nil {
		panic(fmt.Sprintf("read Bls public key of delegate failed, %v", err))
	}
	blsPub, err := bls.PublicKeyFromBytes(blsPubBytes)
	if err != nil {
		panic(fmt.Sprintf("read Bls public key of delegate failed, %v", err))
	}

	return pubKey, &blsPub
}
