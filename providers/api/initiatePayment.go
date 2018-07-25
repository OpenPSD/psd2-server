package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// InitiatePayment returns an payment by ID
func (s Psd2HttpServer) InitiatePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "InitiatePayment with method, %s!\n", ps.ByName("payment-method"))
}
