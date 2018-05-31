// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AvailableAccounts AvailableAccounts
// swagger:model AvailableAccounts
type AvailableAccounts string

const (

	// AvailableAccountsAllAccounts captures enum value "allAccounts"
	AvailableAccountsAllAccounts AvailableAccounts = "allAccounts"

	// AvailableAccountsAllAccountsWithBalances captures enum value "allAccountsWithBalances"
	AvailableAccountsAllAccountsWithBalances AvailableAccounts = "allAccountsWithBalances"
)

// for schema
var availableAccountsEnum []interface{}

func init() {
	var res []AvailableAccounts
	if err := json.Unmarshal([]byte(`["allAccounts","allAccountsWithBalances"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		availableAccountsEnum = append(availableAccountsEnum, v)
	}
}

func (m AvailableAccounts) validateAvailableAccountsEnum(path, location string, value AvailableAccounts) error {
	if err := validate.Enum(path, location, value, availableAccountsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this available accounts
func (m AvailableAccounts) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAvailableAccountsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}