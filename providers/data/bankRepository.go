package data

import "github.com/openpsd/psd2-server/entities"

type BankRepository interface {
	GetAccounts() []entities.AccountDetails
}
