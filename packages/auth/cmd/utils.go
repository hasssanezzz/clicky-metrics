package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var key = []byte(os.Getenv("JWT_SECRET"))

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
		"errors": errors,
		"meta":   nil,
	}

	w.WriteHeader(code)
	WriteJson(w, data)
}

func WriteApiResponse(w http.ResponseWriter, res interface{}) {
	data := map[string]interface{}{
		"data": res,
		"meta": nil,
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

func CreateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * 24 * time.Hour).Unix()
	claims["username"] = username

	return token.SignedString(key)
}

func DecodeToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", fmt.Errorf("can not extract data from token")
	}

	for key, value := range claims {
		if key == "username" {
			username, ok := value.(string)
			if !ok {
				break
			}
			return username, nil
		}
	}

	return "", fmt.Errorf("can not extract data from token")
}
