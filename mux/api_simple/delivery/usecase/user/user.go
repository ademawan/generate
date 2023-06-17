package user

import (
	domain_redis "api_simple/domain/redis"
	response "api_simple/domain/response"
	domain "api_simple/domain/user"
	"time"

	"errors"
	"net/http"
	"net/mail"
	"regexp"
	"strings"
	"unicode"
)

type UserUsecase struct {
	redis domain_redis.RedisRepository
}

func NewUserUsecase(redis domain_redis.RedisRepository) *UserUsecase {
	return &UserUsecase{redis: redis}
}

func (u *UserUsecase) Register(request *domain.UserRegisterRequestFormat, r *http.Request) (*response.Response, int, error) {
	timeNow := time.Now()
	err := u.Validate(request)
	response := &response.Response{}

	if err != nil {
		response.Message = err.Error()
		return response, http.StatusBadRequest, err
	}
	response.Status = true
	response.Message = "SUCCESS REGISTER USER"
	response.Data = domain.UserRegisterResponseFormat{
		Name:      request.Name,
		Address:   request.Address,
		Age:       request.Age,
		Email:     request.Email,
		CreatedAt: timeNow.String(),
		UpdatedAt: timeNow.String(),
	}
	return response, http.StatusOK, nil
}
func (u *UserUsecase) Validate(request *domain.UserRegisterRequestFormat) error {
	var err error
	errorMessage := []string{}
	if request.Name == "" || request.Address == "" || request.Age <= 0 || request.Email == "" || request.Password == "" {

		if request.Name == "" {
			errorMessage = append(errorMessage, "Name is required")
		} else if request.Address == "" {
			errorMessage = append(errorMessage, "Address is required")

		} else if request.Age <= 0 {
			errorMessage = append(errorMessage, "Age is required")

		} else if request.Email == "" {
			errorMessage = append(errorMessage, "Email is required")

		} else {
			errorMessage = append(errorMessage, "Password is required")

		}

		if len(errorMessage) != 0 {
			message := strings.Join(errorMessage, "|")
			err = errors.New(message)
		}
		return err
	}

	_, err = mail.ParseAddress(request.Email)
	if err != nil {
		return errors.New("invalid email address")
	}

	if len(request.Password) < 8 {
		errorMessage = append(errorMessage, "length new_password at least 8 characters")
	}
	if !regexp.MustCompile(`\d`).MatchString(request.Password) {
		errorMessage = append(errorMessage, "new_password must be contain alphanumeric")
	}
	if !hasSymbol(request.Password) {
		errorMessage = append(errorMessage, "new_password must be contain special character")
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
