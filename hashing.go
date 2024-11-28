package idkit_go

import (
	"golang.org/x/crypto/sha3"
	"math/big"
)

func hashToField(input []byte) *big.Int {
	keccakHash := keccak256(input)
	hash := new(big.Int).SetBytes(keccakHash)
	return new(big.Int).Rsh(hash, 8)
}

func keccak256(input []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(input)
	return hash.Sum(nil)
}
