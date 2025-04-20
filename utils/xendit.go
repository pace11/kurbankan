package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type VirtualAccountRequest struct {
	ExternalID     string  `json:"external_id"`
	BankCode       string  `json:"bank_code"`
	Name           string  `json:"name"`
	ExpectedAmount float64 `json:"expected_amount"`
	IsClosed       bool    `json:"is_closed"`
}

func CreateVirtualAccount(externalID, bankCode, name string, amount float64) (map[string]any, error) {
	payload := VirtualAccountRequest{
		ExternalID:     externalID,
		BankCode:       bankCode,
		Name:           name,
		ExpectedAmount: amount,
		IsClosed:       true,
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.xendit.co/callback_virtual_accounts", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv("XENDIT_API_KEY")
	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("xendit error: %v", result)
	}

	return result, nil
}
