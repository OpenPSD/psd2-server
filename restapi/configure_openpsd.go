// This file is safe to edit. Once it exists it will not be overwritten

// /**
//  * OpenPSD PSD2 Server
//  * (C) 2018 by OpenPSD - openpsd.org
//  * Released under the GNU General Public License
//  */

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/openpsd/psd2-server/restapi/operations"
	"github.com/openpsd/psd2-server/restapi/operations/account"
	"github.com/openpsd/psd2-server/restapi/operations/consent"
	"github.com/openpsd/psd2-server/restapi/operations/funds"
	"github.com/openpsd/psd2-server/restapi/operations/payments"
)

//go:generate swagger generate server --target .. --name openpsd --spec ../../../../../../apis/psd2/psd2_swagger2.yaml

func configureFlags(api *operations.OpenpsdAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OpenpsdAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AccountAccountsBalancesByAccountIDGetHandler = account.AccountsBalancesByAccountIDGetHandlerFunc(func(params account.AccountsBalancesByAccountIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation account.AccountsBalancesByAccountIDGet has not yet been implemented")
	})
	api.AccountAccountsByAccountIDGetHandler = account.AccountsByAccountIDGetHandlerFunc(func(params account.AccountsByAccountIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation account.AccountsByAccountIDGet has not yet been implemented")
	})
	api.AccountAccountsGetHandler = account.AccountsGetHandlerFunc(func(params account.AccountsGetParams) middleware.Responder {
		return middleware.NotImplemented("operation account.AccountsGet has not yet been implemented")
	})
	api.AccountAccountsTransactionsByAccountIDAndResourceIDGetHandler = account.AccountsTransactionsByAccountIDAndResourceIDGetHandlerFunc(func(params account.AccountsTransactionsByAccountIDAndResourceIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation account.AccountsTransactionsByAccountIDAndResourceIDGet has not yet been implemented")
	})
	api.AccountAccountsTransactionsByAccountIDGetHandler = account.AccountsTransactionsByAccountIDGetHandlerFunc(func(params account.AccountsTransactionsByAccountIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation account.AccountsTransactionsByAccountIDGet has not yet been implemented")
	})
	api.PaymentsBulkPaymentsByPaymentProductPostHandler = payments.BulkPaymentsByPaymentProductPostHandlerFunc(func(params payments.BulkPaymentsByPaymentProductPostParams) middleware.Responder {
		return middleware.NotImplemented("operation payments.BulkPaymentsByPaymentProductPost has not yet been implemented")
	})
	api.ConsentConsentsByConsentIDDeleteHandler = consent.ConsentsByConsentIDDeleteHandlerFunc(func(params consent.ConsentsByConsentIDDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation consent.ConsentsByConsentIDDelete has not yet been implemented")
	})
	api.ConsentConsentsByConsentIDGetHandler = consent.ConsentsByConsentIDGetHandlerFunc(func(params consent.ConsentsByConsentIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation consent.ConsentsByConsentIDGet has not yet been implemented")
	})
	api.ConsentConsentsByConsentIDPutHandler = consent.ConsentsByConsentIDPutHandlerFunc(func(params consent.ConsentsByConsentIDPutParams) middleware.Responder {
		return middleware.NotImplemented("operation consent.ConsentsByConsentIDPut has not yet been implemented")
	})
	api.ConsentConsentsPostHandler = consent.ConsentsPostHandlerFunc(func(params consent.ConsentsPostParams) middleware.Responder {
		return middleware.NotImplemented("operation consent.ConsentsPost has not yet been implemented")
	})
	api.ConsentConsentsStatusByConsentIDGetHandler = consent.ConsentsStatusByConsentIDGetHandlerFunc(func(params consent.ConsentsStatusByConsentIDGetParams) middleware.Responder {
		return middleware.NotImplemented("operation consent.ConsentsStatusByConsentIDGet has not yet been implemented")
	})
	api.FundsFundsConfirmationsPostHandler = funds.FundsConfirmationsPostHandlerFunc(func(params funds.FundsConfirmationsPostParams) middleware.Responder {
		return middleware.NotImplemented("operation funds.FundsConfirmationsPost has not yet been implemented")
	})
	api.PaymentsPaymentIDByPaymentServiceAndPaymentProductGetHandler = payments.PaymentIDByPaymentServiceAndPaymentProductGetHandlerFunc(func(params payments.PaymentIDByPaymentServiceAndPaymentProductGetParams) middleware.Responder {
		return middleware.NotImplemented("operation payments.PaymentIDByPaymentServiceAndPaymentProductGet has not yet been implemented")
	})
	api.PaymentsPaymentIDByPaymentServiceAndPaymentProductPutHandler = payments.PaymentIDByPaymentServiceAndPaymentProductPutHandlerFunc(func(params payments.PaymentIDByPaymentServiceAndPaymentProductPutParams) middleware.Responder {
		return middleware.NotImplemented("operation payments.PaymentIDByPaymentServiceAndPaymentProductPut has not yet been implemented")
	})
	api.PaymentsPaymentIDStatusByPaymentServiceAndPaymentProductGetHandler = payments.PaymentIDStatusByPaymentServiceAndPaymentProductGetHandlerFunc(func(params payments.PaymentIDStatusByPaymentServiceAndPaymentProductGetParams) middleware.Responder {
		return middleware.NotImplemented("operation payments.PaymentIDStatusByPaymentServiceAndPaymentProductGet has not yet been implemented")
	})
	api.PaymentsPaymentsByPaymentProductPostHandler = payments.PaymentsByPaymentProductPostHandlerFunc(func(params payments.PaymentsByPaymentProductPostParams) middleware.Responder {
		return middleware.NotImplemented("operation payments.PaymentsByPaymentProductPost has not yet been implemented")
	})
	api.PaymentsPeriodicPaymentsByPaymentProductPostHandler = payments.PeriodicPaymentsByPaymentProductPostHandlerFunc(func(params payments.PeriodicPaymentsByPaymentProductPostParams) middleware.Responder {
		return middleware.NotImplemented("operation payments.PeriodicPaymentsByPaymentProductPost has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
