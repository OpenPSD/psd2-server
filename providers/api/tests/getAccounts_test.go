package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/openpsd/psd2-server/providers/api"
	"github.com/openpsd/psd2-server/providers/api/header"
)

func TestGetAccounts(t *testing.T) {
	h, _ := api.NewMockedPsd2HttpServer()
	req, _ := http.NewRequest("GET", "/accounts", nil)
	rr := httptest.NewRecorder()
	req.Header.Set(header.ContentType, header.ApplicationJSON)
	h.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
}
