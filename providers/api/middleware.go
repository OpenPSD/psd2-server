package api

import (
	"net/http"
	"time"

	"github.com/openpsd/psd2-server/providers/api/header"
)

func (s Psd2HttpServer) timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 2*time.Second, "timed out")
}

func (s Psd2HttpServer) checkContentTypeHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get(header.ContentType)
		if contentType == header.ApplicationJSON || contentType == header.ApplicationXML {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
		}
	})
}
