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

	type AuthorisePayment struct {
		PsuData struct {
			Password string `json:"password"`
		} `json:"psuData"`
	}

	var authorisePayment AuthorisePayment

	err := json.NewDecoder(r.Body).Decode(&authorisePayment)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("Authorise Request Body: ", authorisePayment)

	authoriseResponse := `{
		"scaStatus": "psuAuthenticated",
	  "_links":{
	  "authoriseTransaction": {"href": "/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456"}
		}
	  }`
	fmt.Fprintf(w, authoriseResponse)

}
