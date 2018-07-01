package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/openpsd/psd2-server/providers/api/header"
	"github.com/openpsd/psd2-server/providers/data"
	"github.com/openpsd/psd2-server/providers/data/inMemory"
)

type Psd2HttpServer struct {
	bankRepository data.BankRepository
}

func NewMockedPsd2HttpServer() http.Handler {
	inMemBankRepository := inMemory.InMemBankRepository{}
	inMemBankRepository.CreateTestData()
	return Psd2HttpServerFactory(&inMemBankRepository)
}

func Psd2HttpServerFactory(b data.BankRepository) http.Handler {
	s := Psd2HttpServer{
		bankRepository: b,
	}
	routes := s.createRoutes()
	chain := alice.New(s.timeoutHandler, s.checkContentTypeHandler).Then(routes)
	return chain
}

func (s Psd2HttpServer) index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "PSD2 server reference implementation!\n")
}

func (s Psd2HttpServer) timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 2*time.Second, "timed out")
}

func (s Psd2HttpServer) checkContentTypeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get(header.ContentType)
		if contentType == header.ApplicationJSON || contentType == header.ApplicationXML {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
		}
	})
}

func (s Psd2HttpServer) createRoutes() http.Handler {
	routes := httprouter.New()
	routes.GET("/", s.index)
	routes.GET("/accounts/:account-id", s.GetAccountsByID)
	routes.GET("/accounts", s.GetAccounts)

	return routes
}
