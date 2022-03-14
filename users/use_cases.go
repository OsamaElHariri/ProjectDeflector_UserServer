package users

import (
	uuid "github.com/satori/go.uuid"
	"projectdeflector.users/repositories"
)

type UseCase struct {
	Repo repositories.Repository
}

func (useCase UseCase) CreateNewAnonymousUser() (User, error) {
	uuid := uuid.NewV4().String()
	useCase.Repo.InsertUser(uuid)
	return useCase.GetUser(uuid)
}

func (useCase UseCase) UpdateUser(uuid string, nickname string) (User, error) {
	useCase.Repo.UpdateUser(uuid, nickname)
	return useCase.GetUser(uuid)
}

func (useCase UseCase) GetUser(uuid string) (User, error) {
	userResult, err := useCase.Repo.FindUser(uuid)

	if err != nil {
		return User{}, err
	}

	user := User{
		Uuid:     userResult.Uuid,
		Nickname: userResult.Nickname,
	}

	return user, nil
}
