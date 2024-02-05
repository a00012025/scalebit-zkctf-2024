package types

import (
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
)

type Point struct {
	X, Y fr.Element
}

type Message struct {
	Commitment bn254.G1Affine
	Text       string
	Proof      kzg.OpeningProof
}
