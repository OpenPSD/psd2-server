package entities

import "github.com/go-openapi/strfmt"

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
