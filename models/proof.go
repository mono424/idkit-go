package models

type VerificationLevel string

const (
	VerificationLevelOrb    VerificationLevel = "orb"
	VerificationLevelDevice VerificationLevel = "device"
)

type Proof struct {
	Proof             string            `json:"proof"`
	MerkleRoot        string            `json:"merkle_root"`
	NullifierHash     string            `json:"nullifier_hash"`
	VerificationLevel VerificationLevel `json:"verification_level"`
}
