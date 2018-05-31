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

// Protocol protocol
// swagger:model Protocol
type Protocol string

const (

	// ProtocolHTTP captures enum value "http"
	ProtocolHTTP Protocol = "http"

	// ProtocolHTTPS captures enum value "https"
	ProtocolHTTPS Protocol = "https"
)

// for schema
var protocolEnum []interface{}

func init() {
	var res []Protocol
	if err := json.Unmarshal([]byte(`["http","https"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protocolEnum = append(protocolEnum, v)
	}
}

func (m Protocol) validateProtocolEnum(path, location string, value Protocol) error {
	if err := validate.Enum(path, location, value, protocolEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this protocol
func (m Protocol) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProtocolEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}