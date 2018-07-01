package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/openpsd/psd2-server/providers/api"
	"github.com/openpsd/psd2-server/providers/api/header"
)

func TestContentTypeCheck(t *testing.T) {
	h, _ := api.NewMockedPsd2HttpServer()
	req, _ := http.NewRequest("GET", "/accounts", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	t.Logf("test if request status is set to %d if content-type is not checked", http.StatusNotAcceptable)
	if status := rr.Code; status != http.StatusNotAcceptable {
		t.Errorf("want %d got %d", http.StatusNotAcceptable, status)
	}

	t.Logf("test if request status processed if content-type is set correctly")
	req.Header.Set(header.ContentType, header.ApplicationJSON)
	if status := rr.Code; status != http.StatusNotAcceptable {
		t.Errorf("want %d got %d", http.StatusOK, status)
	}
}
