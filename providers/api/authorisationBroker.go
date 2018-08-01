package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// AuthorisationBroker calls authorisation method
func (s Psd2HttpServer) AuthorisationBroker(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	htmlBody, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(htmlBody))
	fmt.Println("Authorisation Broker called with body:", string(htmlBody))

	switch {
	case strings.Contains(string(htmlBody), "scaAuthenticationData"):
		fmt.Println("Found scaAuthenticationData")
		s.AuthoriseTransaction(w, r, ps)

	case strings.Contains(string(htmlBody), "psuData"):
		fmt.Println("Found psuData")
		s.AuthorisePSU(w, r, ps)
	default:
		panic("Authorisation Broker doesn't understand request")

	}

}
