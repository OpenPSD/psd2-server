package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// InitiatePayment returns an payment by ID
func (s Psd2HttpServer) InitiatePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//fmt.Fprintf(w, "InitiatePayment with method: %s \n", ps.ByName("payment-product"))
	//body, err := ioutil.ReadAll(r.Body)
	//decoder := json.NewDecoder(r.Body)

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

	// var requestbody struct {
	// 	// httpbin.org sends back key/value pairs, no map[string][]string
	// 	InstructedAmount struct {
	// 		Currency string `json:"currency"`
	// 		Amount   string `json:"amount"`
	// 	}
	// 	DebtorAccount struct {
	// 		Iban string `json:"iban"`
	// 	}
	// 	CreditorName string `json:"creditorName"`
	// }

	err := json.NewDecoder(r.Body).Decode(&paymentsSepaCT)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(paymentsSepaCT)
	fmt.Println(paymentsSepaCT.InstructedAmount.Amount)

	// if err != nil {
	// 	log.Printf("Error reading body: %v", err)
	// 	http.Error(w, "can't read body", http.StatusBadRequest)
	// 	return
	// } else {
	// 	log.Printf("Body: %v", body)
	// }

	ctResponse := `{
		"transactionStatus": "RCVD",
		"paymentId": "axa-pay-paymentid-1234",
		"_links": {
	"startAuthenticationWithPsuAuthentication": {"href": "/payments/axa-pay-paymentid-1234/authorisations"},
	"self": {"href": "/payments/axa-pay-paymentid-1234"} }
	}`
	fmt.Fprintf(w, ctResponse)

}
