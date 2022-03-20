package repositories

type RepositoryFactory interface {
	GetRepository() (Repository, func(), error)
}

type Repository interface {
	InsertUser(user UserInsertRequest) (string, error)
	FindUserByUuid(uuid string) (UserResult, error)
	FindUser(uuid string) (UserResult, error)
	UpdateUser(id string, user UserUpdateRequest) (UserUpdateRequest, error)
	UpdateUserStats(id string, statUpdate DbGameStat) error
}

func GetRepositoryFactory() RepositoryFactory {
	return getMongoRepositoryFactory()
}
