package user

import (
	response "api_simple/domain/response"
	"net/http"
)

type UserUsecase interface {
	Register(request *UserRegisterRequestFormat, r *http.Request) (*response.Response, int, error)
}
