package user_test

import (
	"errors"
	"testing"
	"unit-test/controllers/user"
	"unit-test/entities"

	"github.com/stretchr/testify/assert"
)

type MockUserSuccess struct{}

func (r *MockUserSuccess) Create(user *entities.User) (*entities.User, error) {
	return &entities.User{Name: "Ade Mawan"}, nil
}
func (r *MockUserSuccess) Update(user *entities.User) (*entities.User, error) {
	return &entities.User{Name: "Ade Mawan"}, nil
}

type MockUserFailed struct{}

func (r *MockUserFailed) Create(user *entities.User) (*entities.User, error) {
	return nil, errors.New("error")
}
func (r *MockUserFailed) Update(user *entities.User) (*entities.User, error) {
	return nil, errors.New("error")
}

func TestLogin(t *testing.T) {
	t.Run("1. Success Create User Test", func(t *testing.T) {
		userController := user.NewUserController(&MockUserSuccess{})
		res, err := userController.Create()
		assert.NotNil(t, res)
		assert.Nil(t, err)

	})
	t.Run("2. Failed Create User Test", func(t *testing.T) {
		userController := user.NewUserController(&MockUserFailed{})
		res, err := userController.Create()
		assert.Nil(t, res)
		assert.NotNil(t, err)

	})
}
