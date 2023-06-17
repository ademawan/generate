package user

import (
	domain "api_simple/domain/user"
	helper "api_simple/helper"
	"encoding/json"

	"net/http"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(u domain.UserUsecase) *UserController {
	return &UserController{u}
}

func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	requestBody := &domain.UserRegisterRequestFormat{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {

		helper.RespondWithJSON(w, http.StatusBadRequest, &helper.ResponseError{
			Message: "Invalid request payload",
			Status:  http.StatusBadRequest,
		})
		return
	}
	defer r.Body.Close()
	res, statusCode, err := c.userUsecase.Register(requestBody, r)
	if err != nil {
		helper.RespondWithJSON(w, statusCode, res)
		return
	}

	helper.RespondWithJSON(w, statusCode, res)

}
