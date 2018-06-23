package models

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
	Resourceid      string     `json:"resourceid,omitempty"`
	Usage           string     `json:"usage,omitempty"`
}
