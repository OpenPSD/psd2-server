package accounts

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Println(data.Bank.AccountProvider.GetAccounts())
	json.NewEncoder(w).Encode(data.Bank.AccountProvider.GetAccounts())
}
