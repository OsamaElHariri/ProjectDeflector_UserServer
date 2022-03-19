package repositories

type RepositoryFactory interface {
	GetRepository() (Repository, func(), error)
}

type Repository interface {
	InsertUser(user UserInsertRequest)
	FindUser(uuid string) (UserResult, error)
	UpdateUser(user UserInsertRequest) (UserInsertRequest, error)
	UpdateUserStats(id string, statUpdate DbGameStat) error
}

func GetRepositoryFactory() RepositoryFactory {
	return getMongoRepositoryFactory()
}
