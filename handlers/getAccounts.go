package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/openpsd/psd2-server/models"
)

// GetAccountsByID return an account by ID
func GetAccountsByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "account, %s!\n", ps.ByName("id"))
}

// GetAccounts returns a list of accounts
func GetAccounts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var accounts []models.AccountDetails
	json.NewEncoder(w).Encode(accounts)
}
