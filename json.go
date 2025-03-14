package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 { // to check if it's a server error
		log.Printf("RespondWithError(w, %d, %s)\n", code, message)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{Error: message})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
