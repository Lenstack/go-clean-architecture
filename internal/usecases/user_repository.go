package usecases

import "github.com/Lenstack/clean-architecture/internal/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
	FindById(string) (domain.User, error)
	Create(domain.User) (interface{}, error)
	Update(string, domain.User) (interface{}, error)
	Delete(string) (interface{}, error)
}
