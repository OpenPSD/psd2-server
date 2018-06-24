package data

import (
	"github.com/openpsd/psd2-server/data/sandbox"
	"github.com/openpsd/psd2-server/models"
)

type AccountProvider interface {
	GetAccounts() []models.AccountDetails
}

type ReferenceBank struct {
	AccountProvider AccountProvider
}

func NewReferenceBank(accountProvider AccountProvider) ReferenceBank {
	rb := ReferenceBank{
		AccountProvider: accountProvider,
	}
	return rb
}

func NewSandboxBank() ReferenceBank {
	sandboxAccountProvider := sandbox.NewSandboxAccountProvider()
	rb := NewReferenceBank(sandboxAccountProvider)
	return rb
}

// Bank to use
var Bank ReferenceBank
