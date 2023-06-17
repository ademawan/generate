package domain

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"unicode"
)

type Response struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

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
type RedisValue struct {
	Increment int    `json:"increment"`
	Email     string `json:"email"`
	UUID      string `json:"uuid"`
}

//================

//===============================

type ResponseRequestReset struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type ResponseFailed400 struct {
	Code    string `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
type ResponseFailedSubmitNewValue400 struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

//

type ResponseValidateRequestReset struct {
	Type         string      `json:"type"`
	Tag          string      `json:"tag"`
	Requirements Requirement `json:"requirements"`
	Token        string      `json:"token"`
}
type Requirement struct {
	Schema      string     `json:"schema"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Required    []string   `json:"required"`
	Properties  Properties `json:"properties"`
}
type Properties struct {
	Details struct {
		Descriptions string `json:"descriptions"`
		Type         string `json:"type"`
	} `json:"details"`
}

type ResponseSubmitNwwValue struct {
	Type      string
	Tag       string
	Status    struct{ Success bool }
	Additions struct{}
}

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("invalid email address")
	}
	return nil
}
func ValidateRequestPayload(event string, payload interface{}) error {
	errorMessage := []string{}
	var err error
	if event == "HTTP|RESET_PASSWORD|VALIDATE_REQUEST_RESET|" {
		p := payload.(RequestValidateReset)
		if p.Code == "" {
			errorMessage = append(errorMessage, "code is required")
		}
		if p.Token == "" {
			errorMessage = append(errorMessage, "token is required")
		}

	} else if event == "HTTP|RESET_PASSWORD|SUBMIT_NEW_VALUE|" {
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
