package idkit_go

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"idkit-go/models"
	"net/http"
	"strings"
)

type IdKit struct {
	apiUrl string
}

const defaultApiUrl = "https://developer.worldcoin.org/api/v2"

func New(config models.Config) *IdKit {
	apiUrl := defaultApiUrl
	if config.ApiUrl != "" {
		apiUrl = strings.TrimRight(config.ApiUrl, "/")
	}
	return &IdKit{apiUrl: apiUrl}
}

func (k *IdKit) postReq(path string, body []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", k.apiUrl, path)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "idkit-go")

	client := &http.Client{}
	return client.Do(req)
}

func (k *IdKit) VerifyProof(proof models.Proof, appID string, action string, signal []byte) error {
	var signalHash *string
	if len(signal) > 0 {
		hash := fmt.Sprintf("0x%x", hashToField(signal))
		signalHash = &hash
	}

	requestBody := models.VerificationRequest{
		Action:            action,
		Proof:             proof.Proof,
		MerkleRoot:        proof.MerkleRoot,
		NullifierHash:     proof.NullifierHash,
		VerificationLevel: string(proof.VerificationLevel),
		SignalHash:        signalHash,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to serialize request body: %w", err)
	}

	resp, err := k.postReq(fmt.Sprintf("verify/%s", appID), body)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		var errorResponse models.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return fmt.Errorf("failed to decode error response: %w", err)
		}
		return fmt.Errorf("verification failed: %s - %s", errorResponse.Code, errorResponse.Detail)
	default:
		return errors.New("unexpected response from server")
	}
}
