package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string, err error) {
	if err != nil {
		log.Printf("Error: %s", err)
	}

	if code > 499 {
		log.Printf("Responding with 5XX error: %s\n", err)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error encoding response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Code: %v\tPayload: %v", code, data)
	w.WriteHeader(code)
	w.Write(data)
}
