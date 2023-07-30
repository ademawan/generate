package user

import "service-example/entities"

type User interface {
	Create(user *entities.User) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
}
