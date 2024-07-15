package cmd

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func WriteJson(w http.ResponseWriter, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

func ReadJson(w http.ResponseWriter, r *http.Request, value interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&value); err != nil {
		http.Error(w, "Coundn't parse request body", http.StatusBadRequest)
		return err
	}

	return nil
}

func WriteApiError(w http.ResponseWriter, code int, errors map[string][]string) {
	data := map[string]interface{}{
		"status": "error",
		"errors": errors,
		"meta":   nil,
	}

	w.WriteHeader(code)
	WriteJson(w, data)
}

func WriteApiResponse(w http.ResponseWriter, res interface{}) {
	data := map[string]interface{}{
		"status": "error",
		"data":   res,
		"meta":   nil,
	}

	WriteJson(w, data)
}

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}
