package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// AuthorisePayment authorises a payment by ID
func (s Psd2HttpServer) AuthorisePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(ps.ByName("paymentid"))
	authoriseResponse := `{
		"scaStatus": "psuAuthenticated",
		"scaMethods": [
		  {
			"authenticationType": "SMS_OTP",
			"authenticationVersion": "string",
			"authenticationMethodId": "myAuthenticationID",
			"name": "SMS OTP on phone +49160 xxxxx 28",
			"explanation": "Detailed information about the SCA method for the PSU."
		  }
		],
		"chosenScaMethod": {
		  "authenticationType": "SMS_OTP",
		  "authenticationVersion": "string",
		  "authenticationMethodId": "myAuthenticationID",
		  "name": "SMS OTP on phone +49160 xxxxx 28",
		  "explanation": "Detailed information about the SCA method for the PSU."
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
		  "scaRedirect": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "scaOAuth": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "updatePsuIdentification": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "startAuthorisationWithPsuAuthentication": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "selectAuthenticationMethod": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "authoriseTransaction": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "scaStatus": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "additionalProp1": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "additionalProp2": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234",
		  "additionalProp3": "/v1/payments/sepa-credit-transfers/axa-pay-paymentid-1234"
		},
		"psuMessage": "string"
	  }`
	fmt.Fprintf(w, authoriseResponse)

}
