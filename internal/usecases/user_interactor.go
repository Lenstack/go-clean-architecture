package usecases

import "github.com/Lenstack/clean-architecture/internal/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) Index() (domain.Users, error) {
	return ui.UserRepository.FindAll()
}

func (ui *UserInteractor) Show(userId string) (domain.User, error) {
	return ui.UserRepository.FindById(userId)
}

func (ui *UserInteractor) Store(user domain.User) (interface{}, error) {
	return ui.UserRepository.Create(user)
}

func (ui *UserInteractor) Update(userId string, user domain.User) (interface{}, error) {
	return ui.UserRepository.Update(userId, user)
}

func (ui *UserInteractor) Destroy(userId string) (interface{}, error) {
	return ui.UserRepository.Delete(userId)
}
