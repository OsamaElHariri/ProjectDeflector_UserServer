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
	nickname := getRandomNickname()
	color := getPlayerColors(uuid, 1)[0]

	dbUser := repositories.DbUser{
		Uuid:     uuid,
		Nickname: nickname,
		Color:    color,
	}
	useCase.Repo.InsertUser(dbUser)
	return useCase.GetUser(uuid)
}

func (useCase UseCase) UpdateUser(uuid string, nickname string, color string) (User, error) {
	dbUser := repositories.DbUser{
		Uuid:     uuid,
		Nickname: nickname,
		Color:    color,
	}
	useCase.Repo.UpdateUser(dbUser)
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
		Color:    userResult.Color,
	}

	return user, nil
}
