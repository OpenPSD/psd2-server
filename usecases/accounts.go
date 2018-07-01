package usecases

import (
	"github.com/openpsd/psd2-server/entities"
	"github.com/openpsd/psd2-server/providers/data"
)

func GetAccounts(bankRepository data.BankRepository) []entities.AccountDetails {
	return bankRepository.GetAccounts()
}
