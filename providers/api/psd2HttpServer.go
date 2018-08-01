package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/openpsd/psd2-server/providers/data"
	"github.com/openpsd/psd2-server/providers/data/inMemory"
)

// Psd2HttpServer holds all http handlers for the PSD2 API
type Psd2HttpServer struct {
	bankRepository data.BankRepository
}

// NewMockedPsd2HttpServer creates an mocked PSD2 API server
func NewMockedPsd2HttpServer() (http.Handler, Psd2HttpServer) {
	inMemBankRepository := inMemory.InMemBankRepository{}
	inMemBankRepository.CreateTestData()
	return Psd2HttpServerFactory(&inMemBankRepository)
}

// Psd2HttpServerFactory injects the required dependencies into the PSD2 API server
func Psd2HttpServerFactory(b data.BankRepository) (http.Handler, Psd2HttpServer) {
	s := Psd2HttpServer{
		bankRepository: b,
	}
	routes := s.createRoutes()
	chain := alice.New(s.timeoutHandler, s.checkContentTypeHandler).Then(routes)
	return chain, s
}

func (s Psd2HttpServer) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "PSD2 server reference implementation!\n")
}

// Needs to be refactored
func (s Psd2HttpServer) createRoutes() http.Handler {
	routes := httprouter.New()
	routes.GET("/", s.index)
	routes.GET("/accounts/:account-id", s.GetAccountsByID)
	routes.GET("/accounts", s.GetAccounts)
	routes.POST("/payments/:paymentid", s.InitiatePayment)
	routes.POST("/payments/:paymentid/authorisations", s.StartAuthorisationWithPsuAuthentication)
	routes.PUT("/payments/:paymentid/authorisations/:authorisationid", s.AuthorisationBroker)

	return routes
}
