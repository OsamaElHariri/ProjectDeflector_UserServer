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

func (useCase UseCase) GetUser(uuid string) (User, error) {
	userResult, err := useCase.Repo.FindUser(uuid)

	if err != nil {
		return User{}, err
	}

	user := User{
		Uuid: userResult.Uuid,
	}

	return user, nil
}
