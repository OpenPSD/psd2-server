// Code generated by go-swagger; DO NOT EDIT.

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AccountsByAccountIDGetHandlerFunc turns a function with the right signature into a accounts by account Id get handler
type AccountsByAccountIDGetHandlerFunc func(AccountsByAccountIDGetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AccountsByAccountIDGetHandlerFunc) Handle(params AccountsByAccountIDGetParams) middleware.Responder {
	return fn(params)
}

// AccountsByAccountIDGetHandler interface for that can handle valid accounts by account Id get params
type AccountsByAccountIDGetHandler interface {
	Handle(AccountsByAccountIDGetParams) middleware.Responder
}

// NewAccountsByAccountIDGet creates a new http.Handler for the accounts by account Id get operation
func NewAccountsByAccountIDGet(ctx *middleware.Context, handler AccountsByAccountIDGetHandler) *AccountsByAccountIDGet {
	return &AccountsByAccountIDGet{Context: ctx, Handler: handler}
}

/*AccountsByAccountIDGet swagger:route GET /accounts/{account-id} account accountsByAccountIdGet

account details

Reads a list of bank accounts, with balances where required. It is assumed that a consent of the PSU to this access is already given and stored on the ASPSP system. The addressed list of accounts depends then on the PSU ID and the stored consent addressed by consentId, respectively the OAuth2 access token.

*/
type AccountsByAccountIDGet struct {
	Context *middleware.Context
	Handler AccountsByAccountIDGetHandler
}

func (o *AccountsByAccountIDGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAccountsByAccountIDGetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}