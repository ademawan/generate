package user

import (
	"unit-test/entities"
)

type UserRepository struct {
}

func NewUserRepository() {

}

func (r *UserRepository) Create(user *entities.User) (*entities.User, error) {

	return user, nil

}

func (r *UserRepository) Update(user *entities.User) (*entities.User, error) {
	return user, nil
}
