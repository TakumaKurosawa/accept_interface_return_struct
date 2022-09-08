package usecase

import "sample/service"

type UserUseCaseIF interface {
	GetUserByID(id string) *service.User
}

type UseCase struct {
	ucIF UserUseCaseIF
}

func New(ucIF UserUseCaseIF) *UseCase {
	return &UseCase{
		ucIF: ucIF,
	}
}

func (u *UseCase) GetUserByID(id string) *service.User {
	return u.ucIF.GetUserByID(id)
}
