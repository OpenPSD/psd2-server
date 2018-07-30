package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// InitiatePayment returns an payment by ID
func (s Psd2HttpServer) InitiatePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	type PaymentsSepaCT struct {
		CreditorAccount struct {
			Iban string `json:"iban"`
		} `json:"creditorAccount"`
		CreditorName  string `json:"creditorName"`
		DebtorAccount struct {
			Iban string `json:"iban"`
		} `json:"debtorAccount"`
		InstructedAmount struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"instructedAmount"`
		RemittanceInformationUnstructured string `json:"remittanceInformationUnstructured"`
	}

	var paymentsSepaCT PaymentsSepaCT

	err := json.NewDecoder(r.Body).Decode(&paymentsSepaCT)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(paymentsSepaCT)

	ctResponse := `{
		"transactionStatus": "RCVD",
		"paymentId": "axa-pay-paymentid-1234",
		"_links": {
	"startAuthenticationWithPsuAuthentication": {"href": "/payments/axa-pay-paymentid-1234/authorisations"},
	"self": {"href": "/payments/axa-pay-paymentid-1234"} }
	}`
	fmt.Fprintf(w, ctResponse)

}
