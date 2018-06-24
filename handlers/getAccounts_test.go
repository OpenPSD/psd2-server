package accounts

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/openpsd/psd2-server/data"
)

func TestGetAccounts(t *testing.T) {
	data.Bank = data.NewSandboxBank()
	routes := httprouter.New()
	routes.GET("/accounts", GetAccounts)
	request, _ := http.NewRequest("GET", "/accounts", nil)
	rr := httptest.NewRecorder()
	routes.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}
