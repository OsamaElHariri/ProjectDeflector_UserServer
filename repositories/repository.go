package repositories

type RepositoryFactory interface {
	GetRepository() (Repository, func(), error)
}

type Repository interface {
	InsertUser(uuid string)
	FindUser(uuid string) (FindUserResult, error)
	UpdateUser(uuid string, nickname string) (UpdateUserResult, error)
}

func GetRepositoryFactory() RepositoryFactory {
	return getMongoRepositoryFactory()
}
