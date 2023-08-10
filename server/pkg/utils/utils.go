package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, statusCode int, msg string) {
	type errResp struct {
		Message string `json:"message"`
	}

	resp := errResp{
		Message: msg,
	}

	RespondWithJSON(w, statusCode, resp)
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, response any) {
	dat, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("unable to marshal response \n %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(dat)
}
