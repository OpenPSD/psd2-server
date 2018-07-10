package controllers

import (
	"github.com/openpsd/psd2-server/entities"
	"github.com/openpsd/psd2-server/usecases"
)

type AccountController struct {
}

func NewAccountController() AccountController {
	return AccountController{}
}

func (ac *AccountController) GetAccounts(accounts []entities.AccountDetails) []entities.AccountDetails {
	return usecases.GetAccounts(accounts)
}
