// Copyright 2018 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nkeys

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
)

// kp is the internal struct for a kepypair using seed.
type kp struct {
	seed []byte
}

var (
	curve = crypto.S256()
	SeedSize = curve.Params().BitSize/8+8
)
type PublicKey []byte
type PrivateKey []byte
const SignatureSize = crypto.SignatureLength

func GenerateKey(raw []byte) (PublicKey, PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(curve, bytes.NewReader(raw))
	if err != nil {
		return nil, nil, err
	}
	return crypto.CompressPubkey(&priv.PublicKey), crypto.FromECDSA(priv), err
}

func Sign(priv PrivateKey, input []byte) ([]byte, error) {
	edcsaPriv, err := crypto.ToECDSA(priv)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(crypto.Keccak256(input), edcsaPriv)
}

func Verify(pub PublicKey, message, sig []byte) bool {
	pubSigned, err :=  crypto.SigToPub(crypto.Keccak256(message), sig)
	if err != nil {
		return false
	}
	public, err := crypto.DecompressPubkey(pub)
	if err != nil {
		return false
	}
	return public.Equal(pubSigned)
	// FIXME: it's quite strange that the below API doesn't work
	//return crypto.VerifySignature(pub, crypto.Keccak256(message), sig)
}

// CreatePair will create a KeyPair based on the rand entropy and a type/prefix byte. rand can be nil.
func CreatePair(prefix PrefixByte) (KeyPair, error) {
	var rawSeed = make([]byte, SeedSize)

	_, err := io.ReadFull(rand.Reader, rawSeed[:])
	if err != nil {
		return nil, err
	}

	seed, err := EncodeSeed(prefix, rawSeed[:])
	if err != nil {
		return nil, err
	}
	return &kp{seed}, nil
}

// rawSeed will return the raw, decoded 64 byte seed.
func (pair *kp) rawSeed() ([]byte, error) {
	_, raw, err := DecodeSeed(pair.seed)
	return raw, err
}

// keys will return a 32 byte public key and a 64 byte private key utilizing the seed.
func (pair *kp) keys() (PublicKey, PrivateKey, error) {
	raw, err := pair.rawSeed()
	if err != nil {
		return nil, nil, err
	}
	return GenerateKey(raw)
}

// Wipe will randomize the contents of the seed key
func (pair *kp) Wipe() {
	io.ReadFull(rand.Reader, pair.seed)
	pair.seed = nil
}

// Seed will return the encoded seed.
func (pair *kp) Seed() ([]byte, error) {
	return pair.seed, nil
}

// PublicKey will return the encoded public key associated with the KeyPair.
// All KeyPairs have a public key.
func (pair *kp) PublicKey() (string, error) {
	public, raw, err := DecodeSeed(pair.seed)
	if err != nil {
		return "", err
	}
	pub, _, err := GenerateKey(raw)
	if err != nil {
		return "", err
	}
	pk, err := Encode(public, pub)
	if err != nil {
		return "", err
	}
	return string(pk), nil
}

// PrivateKey will return the encoded private key for KeyPair.
func (pair *kp) PrivateKey() ([]byte, error) {
	_, priv, err := pair.keys()
	if err != nil {
		return nil, err
	}
	return Encode(PrefixBytePrivate, priv)
}

// Sign will sign the input with KeyPair's private key.
func (pair *kp) Sign(input []byte) ([]byte, error) {
	_, priv, err := pair.keys()
	if err != nil {
		return nil, err
	}
	return Sign(priv, input)
}

// Verify will verify the input against a signature utilizing the public key.
func (pair *kp) Verify(input []byte, sig []byte) error {
	pub, _, err := pair.keys()
	if err != nil {
		return err
	}
	if !Verify(pub, input, sig) {
		return ErrInvalidSignature
	}
	return nil
}
