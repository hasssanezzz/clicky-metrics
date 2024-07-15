package cmd

import (
	"fmt"
	"net/http"

	"github.com/hasssanezzz/multi-service-shortner/storage"
)

var UserRepo = storage.UserRepo{}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotFound)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	type RegisterRequestBody struct {
		Email    string `json:"email" validate:"required,email"`
		Username string `json:"username" validate:"required,username"`
		Password string `json:"password" validate:"required,password"`
	}

	// parse request body
	var body RegisterRequestBody
	err := ReadJson(w, r, &body)
	if err != nil {
		return
	}

	// validate body
	if errors, ok := ValidateBody(body); !ok {
		WriteApiError(w, http.StatusBadRequest, errors)
		return
	}

	_, usernameErr := UserRepo.FindByUsername(body.Username)
	_, emailErr := UserRepo.FindByUsername(body.Username)

	if usernameErr == nil || emailErr == nil {
		errors := make(map[string][]string)

		if usernameErr == nil {
			errors["username"] = []string{"Username is in use."}
		}
		if emailErr == nil {
			errors["email"] = []string{"E-mail is in use."}
		}

		WriteApiError(w, http.StatusBadRequest, errors)
		return
	}

	hash, err := HashPassword(body.Password)
	if err != nil {
		WriteApiError(w, http.StatusBadRequest, map[string][]string{
			"password": {"internal server error, please choose another password."},
		})
		return
	}

	newUser := &storage.User{
		Email:    body.Email,
		Username: body.Username,
		Password: string(hash),
	}

	newUser, err = UserRepo.Create(newUser)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	WriteApiResponse(w, newUser)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	type LoginRequestBody struct {
		Username string `json:"username" validate:"required,username"`
		Password string `json:"password" validate:"required,password"`
	}

	// parse request body
	var body LoginRequestBody
	err := ReadJson(w, r, &body)
	if err != nil {
		return
	}

	// validate body
	if errors, ok := ValidateBody(body); !ok {
		WriteApiError(w, http.StatusBadRequest, errors)
		return
	}

	// find user by username
	user, err := UserRepo.FindByUsername(body.Username)
	if err != nil {
		WriteApiError(w, http.StatusBadRequest, map[string][]string{
			"username": {"username does not exist."},
		})
		return
	}

	// match passwords
	if err := UserRepo.VerifyPassword(user.Password, body.Password); err != nil {
		WriteApiError(w, http.StatusBadRequest, map[string][]string{
			"password": {"incorrect username/password combination."},
		})
		return
	}

	WriteJson(w, "Hello world")
}
