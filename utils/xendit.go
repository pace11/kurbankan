package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type VirtualAccountRequest struct {
	ExternalID     string  `json:"external_id"`
	BankCode       string  `json:"bank_code"`
	Name           string  `json:"name"`
	ExpectedAmount float64 `json:"expected_amount"`
	IsClosed       bool    `json:"is_closed"`
	IsSingleUse    bool    `json:"is_single_use"`   // optional
	ExpirationDate string  `json:"expiration_date"` // optional ISO8601
}

// response bisa kamu ubah jadi struct jika ingin lebih strict
func CreateVirtualAccount(externalID, bankCode, name string, amount float64) (map[string]any, error) {
	payload := VirtualAccountRequest{
		ExternalID:     externalID,
		BankCode:       bankCode,
		Name:           name,
		ExpectedAmount: amount,
		IsClosed:       true,
		IsSingleUse:    true,
		ExpirationDate: time.Now().Add(48 * time.Hour).UTC().Format(time.RFC3339), // optional
	}

	body, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://api.xendit.co/callback_virtual_accounts", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(os.Getenv("XENDIT_API_KEY"), "")
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
