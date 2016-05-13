package handlers

import (
	"net/http"

	"encoding/json"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal("Hello world!")

	w.Write(j)
}
