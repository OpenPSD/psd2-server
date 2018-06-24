package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/openpsd/psd2-server/data"
)

// GetAccountsByID return an account by ID
func GetAccountsByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "account, %s!\n", ps.ByName("id"))
}

// GetAccounts returns a list of accounts
func GetAccounts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json.NewEncoder(w).Encode(data.Bank.AccountProvider.GetAccounts())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
