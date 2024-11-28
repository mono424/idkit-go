# `idkit_go` - Go SDK for Worldcoin's IDKit

`idkit_go` is a Go SDK for integrating with Worldcoin's IDKit, a service that allows you to verify user identities using World ID. This SDK provides the necessary tools to interact with the IDKit API and verify user proofs securely and easily in your Go applications.

## Installation

You can install the `idkit_go` package via `go get`:

```bash
go get github.com/mono424/idkit-go
```

## Features

**🛡️Verify Proof**: Verifies the Zero-knowledge proof using the provided proof, appID, and signal.

## How to use

```go
import (
	"fmt"
	"idkit-go/models"
	"idkit-go"
)

func main() {
	// Define the proof to be verified
	proof := models.Proof{
		Proof:         "proof_string",
		MerkleRoot:    "merkle_root_string",
		NullifierHash: "nullifier_hash_string",
		VerificationLevel: "orb",
	}

	// Define the appID and action
	appID := "app_12345"
	action := "sign_in"

	// Define an optional signal (empty byte slice if not used)
	signal := []byte("signal_data")

	// Create a new IDKit client
	client := idkit_go.New(models.Config{})

	// Verify the proof
	err := client.VerifyProof(proof, appID, action, signal)
	if err != nil {
		fmt.Printf("Verification failed: %v\n", err)
		return
	}

	fmt.Println("Verification successful!")
}
```