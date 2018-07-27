package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AuthorisePayment authorises a payment by ID
func (s Psd2HttpServer) AuthorisePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Authorise payment with payment ID:", ps.ByName("paymentid"))
	fmt.Println("Authorise payment with authorisation ID:", ps.ByName("authorisationid"))

	authoriseResponse := `{
		"scaStatus": "psuAuthenticated",
	  "_links":{
	  "authoriseTransaction": {"href": "/v1/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456"}
		}
	  }`
	fmt.Fprintf(w, authoriseResponse)

}
