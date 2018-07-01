package entities

type Links struct {
	Account                    string `json:"account,omitempty"`
	AuthoriseTransaction       string `json:"authoriseTransaction,omitempty"`
	Balances                   string `json:"balances,omitempty"`
	Download                   string `json:"download,omitempty"`
	First                      string `json:"first,omitempty"`
	Last                       string `json:"last,omitempty"`
	Next                       string `json:"next,omitempty"`
	Previous                   string `json:"previous,omitempty"`
	ScaOAuth                   string `json:"scaOAuth,omitempty"`
	ScaRedirect                string `json:"scaRedirect,omitempty"`
	SelectAuthenticationMethod string `json:"selectAuthenticationMethod,omitempty"`
	Self                       string `json:"self,omitempty"`
	Status                     string `json:"status,omitempty"`
	Transactions               string `json:"transactions,omitempty"`
	TransactionsDetails        string `json:"transactionsDetails,omitempty"`
	UpdateProprietaryData      string `json:"updateProprietaryData,omitempty"`
	UpdatePsuAuthentication    string `json:"updatePsuAuthentication,omitempty"`
	UpdatePsuIdentification    string `json:"updatePsuIdentification,omitempty"`
}
