package main

import (
	"log"
	"net/http"

	"github.com/openpsd/psd2-server/providers/api"
)

func main() {
	log.Println("start OpenPSD DeliveryThinking PSD2 reference implementation server")
	psd2Api, _ := api.NewMockedPsd2HttpServer()
	err := http.ListenAndServe(":8000", psd2Api)
	if err != nil {
		panic(err.Error())
	}
}
