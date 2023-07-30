package user

import (
	"fmt"
	"service-example/entities"
	"service-example/repository/user"
)

type UserController struct {
	user user.User
}

func NewUserController(u user.User) *UserController {
	return &UserController{user: u}
}
func (c *UserController) Create() (*entities.User, error) {

	res, err := c.user.Create(&entities.User{Name: "Ade Mawan", Age: 29})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(res)
	return res, nil

}
func (c *UserController) Update() (*entities.User, error) {

	res, err := c.user.Update(&entities.User{Name: "Ade Mawan", Age: 29})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(res)
	return res, nil

}
