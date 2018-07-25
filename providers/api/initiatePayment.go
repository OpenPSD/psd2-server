package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// InitiatePayment returns an payment by ID
func (s Psd2HttpServer) InitiatePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//fmt.Fprintf(w, "InitiatePayment with method: %s \n", ps.ByName("payment-product"))
	ctResponse := `{
		"transactionStatus": "ACTC",
		"paymentId": "0007.01",
		"transactionFees": {
		  "currency": "EUR",
		  "amount": "0.42"
		},
		"transactionFeeIndicator": true,
		"scaMethods": [
		  {
			"authenticationType": "SMS_OTP",
			"authenticationVersion": "string",
			"authenticationMethodId": "string",
			"name": "string",
			"explanation": "string"
		  }
		],
		"chosenScaMethod": {
		  "authenticationType": "SMS_OTP",
		  "authenticationVersion": "string",
		  "authenticationMethodId": "string",
		  "name": "string",
		  "explanation": "string"
		},
		"challengeData": {
		  "image": "string",
		  "data": "string",
		  "imageLink": "string",
		  "otpMaxLength": 0,
		  "otpFormat": "characters",
		  "additionalInformation": "string"
		},
		"_links": {
		  "scaRedirect": "string",
		  "scaOAuth": "string",
		  "updatePsuIdentification": "string",
		  "updateProprietaryData": "string",
		  "updatePsuAuthentication": "string",
		  "selectAuthenticationMethod": "string",
		  "authoriseTransaction": "string",
		  "self": "string",
		  "status": "string",
		  "account": "string",
		  "balances": "string",
		  "transactions": "string",
		  "transactionsDetails": "string",
		  "first": "string",
		  "next": "string",
		  "previous": "string",
		  "last": "string",
		  "download": "string"
		},
		"psuMessage": "string",
		"tppMessages": [
		  {
			"category": "ERROR",
			"code": "TOKEN_INVALID",
			"text": "additional text information of the ASPSP up to 512 characters",
			"path": "string"
		  }
		]
	  }`
	fmt.Fprintf(w, ctResponse)

}
