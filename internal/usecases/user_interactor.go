package usecases

import (
	"github.com/Lenstack/clean-architecture/internal/domain/model"
)

type UserInteractor struct {
	UserRepository UserService
}

func (ui *UserInteractor) Index() (model.Users, error) {
	return ui.UserRepository.FindAll()
}

func (ui *UserInteractor) Show(userId string) (model.User, error) {
	return ui.UserRepository.FindById(userId)
}

func (ui *UserInteractor) Store(user model.User) (interface{}, error) {
	return ui.UserRepository.Create(user)
}

func (ui *UserInteractor) Update(userId string, user model.User) (interface{}, error) {
	return ui.UserRepository.Update(userId, user)
}

func (ui *UserInteractor) Destroy(userId string) (interface{}, error) {
	return ui.UserRepository.Delete(userId)
}
