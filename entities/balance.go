package entities

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
)

type BalanceType string

const (
	BalanceTypeClosingBooked    BalanceType = "closingBooked"
	BalanceTypeExpected         BalanceType = "expected"
	BalanceTypeAuthorised       BalanceType = "authorised"
	BalanceTypeOpeningBooked    BalanceType = "openingBooked"
	BalanceTypeInterimAvailable BalanceType = "interimAvailable"
	BalanceTypeForwardAvailable BalanceType = "forwardAvailable"
)

type Balance struct {
	BalanceAmount            float32         `json:"balanceAmount"`
	BalanceType              BalanceType     `json:"balanceType"`
	LastChangeDateTime       strfmt.DateTime `json:"lastChangeDateTime,omitempty"`
	LastCommittedTransaction string          `json:"lastCommittedTransaction,omitempty"`
	ReferenceDate            strfmt.Date     `json:"referenceDate,omitempty"`
}

// Marshal interface implementation
func (m *Balance) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *Balance) Unmarshal(b []byte) error {
	var res Balance
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
