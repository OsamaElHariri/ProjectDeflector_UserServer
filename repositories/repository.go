package repositories

type RepositoryFactory interface {
	GetRepository() (Repository, func(), error)
}

type Repository interface {
	InsertUser(user DbUser)
	FindUser(uuid string) (DbUser, error)
	UpdateUser(user DbUser) (DbUser, error)
}

func GetRepositoryFactory() RepositoryFactory {
	return getMongoRepositoryFactory()
}
