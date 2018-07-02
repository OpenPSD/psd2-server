package entities

import "encoding/json"

type AccountDetails struct {
	Links           Links      `json:"_links,omitempty"`
	Balances        []*Balance `json:"balances"`
	Bban            string     `json:"bban,omitempty"`
	Bic             string     `json:"bic,omitempty"`
	CashAccountType string     `json:"cashAccountType,omitempty"`
	Currency        *string    `json:"currency"`
	Details         string     `json:"details,omitempty"`
	Iban            string     `json:"iban,omitempty"`
	LinkedAccounts  string     `json:"linkedAccounts,omitempty"`
	MaskedPan       string     `json:"maskedPan,omitempty"`
	Msisdn          string     `json:"msisdn,omitempty"`
	Name            string     `json:"name,omitempty"`
	Pan             string     `json:"pan,omitempty"`
	Product         string     `json:"product,omitempty"`
	ResourceID      string     `json:"resourceid,omitempty"`
	Usage           string     `json:"usage,omitempty"`
}

func NewAccountDetails(currency string) AccountDetails {
	c := &currency
	return AccountDetails{
		Currency: c,
	}
}

// Marshal interface implementation
func (m *AccountDetails) Marshal() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return json.Marshal(m)
}

// Unmarshal interface implementation
func (m *AccountDetails) Unmarshal(b []byte) error {
	var res AccountDetails
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
