package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorObject struct {
	Message string `json:"message"`
}

func SendJsonBadRequest(rw http.ResponseWriter, message string) {
	sendJson(rw, http.StatusBadRequest, &errorObject{Message: message})
}

func SendJsonInternalError(rw http.ResponseWriter) {
	sendJson(rw, http.StatusBadRequest, &errorObject{Message: "server internal error"})
}

func SendJsonOk(rw http.ResponseWriter, res interface{}) {
	sendJson(rw, http.StatusOK, res)
}

func sendJson(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
