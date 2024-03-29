package users

import (
	uuid "github.com/satori/go.uuid"
	"projectdeflector.users/repositories"
)

type UseCase struct {
	Repo repositories.Repository
}

func (useCase UseCase) CreateNewAnonymousUser() (string, User, error) {
	uuid := uuid.NewV4().String()
	randomUser := getRandomUser(uuid)

	dbUser := repositories.UserInsertRequest{
		Uuid:     uuid,
		Nickname: randomUser.Nickname,
		Color:    randomUser.Color,
	}
	id, err := useCase.Repo.InsertUser(dbUser)
	if err != nil {
		return "", User{}, err
	}

	user, err := useCase.GetUser(id)

	return uuid, user, err
}

func getRandomUser(uuid string) User {
	return User{
		Nickname: getRandomNickname(),
		Color:    getPlayerColors(uuid, 1)[0],
	}
}

func (useCase UseCase) UpdateUser(id string, nickname string, color string) (User, error) {
	dbUser := repositories.UserUpdateRequest{
		Nickname: nickname,
		Color:    color,
	}
	useCase.Repo.UpdateUser(id, dbUser)
	return useCase.GetUser(id)
}

func (useCase UseCase) GetUser(id string) (User, error) {
	if id == "system" {
		randomUser := getRandomUser(id)
		randomUser.Id = id
		return randomUser, nil
	}

	userResult, err := useCase.Repo.FindUser(id)

	if err != nil {
		return User{}, err
	}

	user := User{
		Id:       userResult.Id,
		Nickname: userResult.Nickname,
		Color:    userResult.Color,
		GameStats: GameStats{
			Games: userResult.GameStats.Games,
			Wins:  userResult.GameStats.Wins,
		},
	}

	return user, nil
}

type GameStatUpdate struct {
	PlayerId string
	Games    int
	Wins     int
}

func (useCase UseCase) UpdateUserStats(updates []GameStatUpdate) {
	for i := 0; i < len(updates); i++ {
		update := repositories.DbGameStat{
			Games: updates[i].Games,
			Wins:  updates[i].Wins,
		}
		useCase.Repo.UpdateUserStats(updates[i].PlayerId, update)
	}
}

func (useCase UseCase) GetAccessToken(uuid string) (string, error) {
	user, err := useCase.Repo.FindUserByUuid(uuid)
	if err != nil {
		return "", err
	}
	return issueJwt(User{
		Id: user.Id,
	})
}

func (useCase UseCase) ValidateAccessToken(token string) (string, error) {
	info, err := validateJwt(token)
	if err != nil {
		return "", err
	}
	return info.UserId, nil
}

func (useCase UseCase) GetUserColors(id string) ([]string, error) {
	user, err := useCase.Repo.FindUser(id)
	if err != nil {
		return nil, err
	}
	colors := getPlayerColors(user.Uuid, 4)
	return colors, nil
}
