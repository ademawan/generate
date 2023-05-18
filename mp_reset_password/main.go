package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"
)

type RequestResetPassword struct {
	Email string `json:"email"`
}

type RequestValidateReset struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}
type RequestSubmitNewValue struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func main() {
	// err := ValidateRequestPayload("VALIDATE_REQUEST_RESET|", RequestValidateReset{Code: "dsjhkjewueaFk", Token: "jhasdasks"})
	err := ValidateRequestPayload("SUBMIT_NEW_VALUE|", RequestSubmitNewValue{NewPassword: "#Adol1"})

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("success")

}

func ValidateRequestPayload(event string, payload interface{}) error {
	errorMessage := []string{}
	var err error
	if event == "VALIDATE_REQUEST_RESET|" {
		p := payload.(RequestValidateReset)
		if p.Code == "" {
			errorMessage = append(errorMessage, "code is required")
		}
		if p.Token == "" {
			errorMessage = append(errorMessage, "token is required")
		}

	} else if event == "SUBMIT_NEW_VALUE|" {
		p := payload.(RequestSubmitNewValue)
		if p.NewPassword == "" {
			errorMessage = append(errorMessage, "new_password is required")
		}
		if len(p.NewPassword) < 8 {
			errorMessage = append(errorMessage, "length new_password at least 8 characters")

		}
		if !regexp.MustCompile(`\d`).MatchString(p.NewPassword) {
			errorMessage = append(errorMessage, "new_password must be contain alphanumeric")
		}
		if !hasSymbol(p.NewPassword) {
			errorMessage = append(errorMessage, "new_password must be contain special character")
		}
		if p.Token == "" {
			errorMessage = append(errorMessage, "token is required")
		}

	}

	if len(errorMessage) != 0 {
		message := strings.Join(errorMessage, "|")
		err = errors.New(message)
	}
	return err
}
func hasSymbol(str string) bool {
	for _, letter := range str {
		if unicode.IsSymbol(letter) || unicode.IsPunct(letter) {
			return true
		}

	}
	return false
}
