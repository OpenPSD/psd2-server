package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/openpsd/psd2-server/handlers"
)

func timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 2*time.Second, "timed out")
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "PSD2 server reference implementation!\n")
}

func main() {
	log.Println("start PSD2 reference implementation server")
	routes := createRoutes()

	chain := alice.New(timeoutHandler).Then(routes)
	http.ListenAndServe(":8000", chain)
}

func createRoutes() http.Handler {
	routes := httprouter.New()
	routes.GET("/", index)
	routes.GET("/accounts/:account-id", accounts.GetAccountsByID)
	routes.GET("/accounts", accounts.GetAccounts)

	return routes
}
