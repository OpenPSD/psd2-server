package sandbox

import "github.com/openpsd/psd2-server/models"

type SandboxAccountProvider struct {
	accounts []models.AccountDetails
}

func (ap SandboxAccountProvider) GetAccounts() []models.AccountDetails {
	return ap.accounts
}

func NewSandboxAccountProvider() SandboxAccountProvider {
	newAccounts := createAccounts()
	sandboxAccountProvider := SandboxAccountProvider{
		accounts: newAccounts,
	}
	return sandboxAccountProvider
}

func createAccounts() []models.AccountDetails {
	var accounts []models.AccountDetails

	a := models.NewAccountDetails("EUR")
	a.Name = "DemoAccount1"
	accounts = append(accounts, a)

	return accounts
}
