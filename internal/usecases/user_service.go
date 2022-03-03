package usecases

import (
	"github.com/Lenstack/clean-architecture/internal/domain/model"
)

type UserService interface {
	FindAll() (model.Users, error)
	FindById(string) (model.User, error)
	Create(model.User) (interface{}, error)
	Update(string, model.User) (interface{}, error)
	Delete(string) (interface{}, error)
}
