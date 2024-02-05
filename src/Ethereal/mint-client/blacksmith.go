package main

import (
	"crypto"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
	"github.com/davecgh/go-spew/spew"
)

func ForgeSword() []fr.Element {
	f := make([]fr.Element, 17)
	h := crypto.SHA256.New()
	ts := time.Now().UnixNano()
	tsB := make([]byte, 8)
	binary.LittleEndian.PutUint64(tsB, uint64(ts))
	h.Write(tsB)
	h.Write([]byte("Soul of a Hero"))
	seed := h.Sum(nil)
	for i := 0; i < 17; i++ {
		f[i].SetBytes(seed[:32])
		h.Reset()
		h.Write(seed)
		seed = h.Sum(nil)
	}
	return f
}

type KeyPairProof struct {
	H              bn254.G1Affine
	PrivateKey     *fr.Element
	PublicKeyG1Aff bn254.G1Affine
}

func CraftBladeSignature(poly []fr.Element, srs *kzg.SRS) (kzg.Digest, *KeyPairProof) {
	//commit the polynomial
	commitment, err := kzg.Commit(poly, srs.Pk)
	if err != nil {
		panic(err)
	}
	fmt.Printf("commitment: \nX: %02x\nY: %02x\n", commitment.X, commitment.Y)

	// compute opening proof at a random point
	var point fr.Element
	point.SetInt64(0)
	proof, err := kzg.Open(poly, point, srs.Pk)
	if err != nil {
		panic(err)
	}

	// claimed value is private key.
	// derive public key from private key
	privateKey := new(big.Int)
	proof.ClaimedValue.BigInt(privateKey)
	fmt.Printf("private key: %s\n", privateKey.String())

	publicKey := new(bn254.G1Affine)
	publicKey.ScalarMultiplication(&srs.Pk.G1[0], privateKey)
	spew.Dump(privateKey, publicKey)
	publicKeyG2 := new(bn254.G2Affine)
	publicKeyG2.ScalarMultiplication(&srs.Vk.G2[0], privateKey)

	pubKeyProof := new(KeyPairProof)
	pubKeyProof.PrivateKey = &proof.ClaimedValue
	pubKeyProof.H = proof.H
	pubKeyProof.PublicKeyG1Aff = *publicKey
	return commitment, pubKeyProof
}
