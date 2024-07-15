package cmd

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate validator.Validate = *validator.New()

func InitValidator() {

	if err := Validate.RegisterValidation("username", validateUsername); err != nil {
		log.Fatalf("error registering username validator: %v", err)
	}

	if err := Validate.RegisterValidation("password", validatePassword); err != nil {
		log.Fatalf("error registering password validator: %v", err)
	}
}

func validateUsername(fl validator.FieldLevel) bool {
	pattern := `^[a-zA-Z0-9_]{4,20}$`
	username := fl.Field().String()
	matched, _ := regexp.MatchString(pattern, username)
	return matched
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	length := len(password)
	return length < 72 && length >= 8
}

func ValidateBody(i interface{}) (map[string][]string, bool) {
	errors := make(map[string][]string)

	if err := Validate.Struct(i); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				errors["email"] = []string{"Invalid E-mail format."}
			case "username":
				errors["username"] = []string{"Usernames can only consist of letters, digits and _, length should be from 4 to 20 characters."}
			case "password":
				errors["password"] = []string{"Password length should be from 8 to 72 characters."}
			case "required":
				field := strings.ToLower(err.Field())
				errors[field] = []string{fmt.Sprintf("%s is requried.", field)}
			}
		}
		return errors, false
	}
	return nil, true
}
