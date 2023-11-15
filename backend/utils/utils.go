package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func HandleCORS(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if (*r).Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return true
	} else {
		return false
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func DecodeJSON(w http.ResponseWriter, r *http.Request, data any) error {
	defer r.Body.Close()

	if json.NewDecoder(r.Body).Decode(&data) != nil {
		fmt.Println("Invalid JSON recieved")
		return errors.New("invalid json recieved")
	}

	return nil
}
