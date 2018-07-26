package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// InitiatePayment returns an payment by ID
func (s Psd2HttpServer) InitiatePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//fmt.Fprintf(w, "InitiatePayment with method: %s \n", ps.ByName("payment-product"))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	} else {
		log.Printf("Body: %v", body)
	}

	ctResponse := `{
		"transactionStatus": "RCVD",
		"paymentId": "axa-pay-paymentid-1234",
		"_links": {
	"startAuthenticationWithPsuAuthentication": {"href": "/payments/axa-pay-paymentid-1234/authorisations"},
	"self": {"href": "/payments/axa-pay-paymentid-1234"} }
	}`
	fmt.Fprintf(w, ctResponse)

}
