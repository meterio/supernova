// Copyright (c) 2020 The Meter.io developers
// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying

// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/blst"
)

type KeyLoader struct {
	baseDir  string
	keysPath string
}

type KeysContent struct {
	Secret string `json:"secret"`
	Pubkey string `json:"pubkey"`
}

func NewKeyLoader(baseDir string) *KeyLoader {
	keysPath := filepath.Join(baseDir, "keys.json")

	return &KeyLoader{
		baseDir:  baseDir,
		keysPath: keysPath,
	}
}

func (k *KeyLoader) Load() (*BlsMaster, error) {
	var secret bls.SecretKey
	var pubkey bls.PublicKey

	var keysContent KeysContent
	if !common.FileExist(k.keysPath) {
		secretKey, err := blst.RandKey()
		if err != nil {
			return nil, err
		}
		keysContent := KeysContent{
			Secret: base64.StdEncoding.EncodeToString(secretKey.Marshal()),
			Pubkey: base64.StdEncoding.EncodeToString(secretKey.PublicKey().Marshal()),
		}
		keysBytes, err := json.Marshal(keysContent)
		if err != nil {
			return nil, err
		}
		fmt.Println("output to ", k.keysPath)
		err = os.WriteFile(k.keysPath, keysBytes, 0600)
		if err != nil {
			return nil, err
		}
	}
	keysBytes, err := os.ReadFile(k.keysPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(keysBytes, &keysContent)
	if err != nil {
		return nil, err
	}
	secretBytes, err := base64.StdEncoding.DecodeString(keysContent.Secret)
	if err != nil {
		return nil, err
	}
	secret, err = bls.SecretKeyFromBytes(secretBytes)

	pubBytes, err := base64.StdEncoding.DecodeString(keysContent.Pubkey)
	if err != nil {
		return nil, err
	}
	pubkey, err = bls.PublicKeyFromBytes(pubBytes)
	return NewBlsMaster(secret, pubkey), nil
}
