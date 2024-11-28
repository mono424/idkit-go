package models

type VerificationRequest struct {
	Action            string  `json:"action"`
	Proof             string  `json:"proof"`
	MerkleRoot        string  `json:"merkle_root"`
	NullifierHash     string  `json:"nullifier_hash"`
	VerificationLevel string  `json:"verification_level"`
	SignalHash        *string `json:"signal_hash,omitempty"`
}
