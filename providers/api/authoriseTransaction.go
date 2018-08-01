package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AuthoriseTransaction authorises a payment by payment- and authorisation ID
func (s Psd2HttpServer) AuthoriseTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Authorise transaction with payment ID:", ps.ByName("paymentid"))
	fmt.Println("Authorise transaction with authorisation ID:", ps.ByName("authorisationid"))

	type AuthoriseTransaction struct {
		ScaAuthenticationData string `json:"scaAuthenticationData"`
	}

	var authoriseTransaction AuthoriseTransaction

	err := json.NewDecoder(r.Body).Decode(&authoriseTransaction)
	if err != nil {
		http.Error(w, err.Error(), 400)
		fmt.Println("Authorise Transaction Request Error: ", err)
		return
	}
	fmt.Println("Authorise Transaction Request Body: ", authoriseTransaction)

	authoriseResponse := `{
		"scaStatus": "finalised",
		"_links":{
	  "scaStatus": {"href":"/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456"}
	  } }
		`
	fmt.Fprintf(w, authoriseResponse)

}
