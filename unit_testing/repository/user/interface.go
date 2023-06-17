package user

import "unit-test/entities"

type User interface {
	Create(user *entities.User) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
}
