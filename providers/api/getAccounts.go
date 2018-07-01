package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/openpsd/psd2-server/providers/api/header"
	"github.com/openpsd/psd2-server/usecases"
)

// GetAccountsByID return an account by ID
func (s Psd2HttpServer) GetAccountsByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "account, %s!\n", ps.ByName("id"))
}

// GetAccounts returns a list of accounts
func (s Psd2HttpServer) GetAccounts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json.NewEncoder(w).Encode(usecases.GetAccounts(s.bankRepository))
	w.Header().Set(header.ContentType, header.ApplicationJSON)
	w.WriteHeader(http.StatusOK)
}
