package inMemory

import (
	"github.com/openpsd/psd2-server/entities"
)

// InMemBankRepository mocks the access to a core banking system
type InMemBankRepository struct {
	accounts []entities.AccountDetails
}

// GetAccounts returns a list of accounts
func (r *InMemBankRepository) GetAccounts() []entities.AccountDetails {
	return r.accounts
}

func (r *InMemBankRepository) CreateTestData() {
	var ac []entities.AccountDetails

	a := entities.NewAccountDetails("EUR")
	a.Name = "DemoAccount1"
	ac = append(ac, a)

	b := entities.NewAccountDetails("CHF")
	b.Name = "DemoAccount2"
	ac = append(ac, b)

	r.accounts = ac
}
