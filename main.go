package main

import (
	"log"
	"net/http"
	"os"

	. "github.com/GermanMontejo/gohcr/utils"
)

func main() {
	address := os.Args[2]
	method := os.Args[1]
	body := CreateRequestBody(method, os.Args)
	req := CreateNewRequest(method, address, body)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Error sending request to server:", err)
		return
	}

	DisplayResponseDetails(resp)
}