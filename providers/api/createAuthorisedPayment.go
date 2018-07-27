package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CreateAuthorisedPayment creates authorisation for a payment by ID
func (s Psd2HttpServer) CreateAuthorisedPayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Create authorisation payment with ID:", ps.ByName("paymentid"))

	authoriseResponse := `{
		"scaStatus": "received",
"psuMessage": "Please use your BankApp for transaction Authorisation.", "_links":{
"scaStatus": {"href":"/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456"}
} }`
	fmt.Fprintf(w, authoriseResponse)

}
