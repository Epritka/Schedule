package interfaces

import "user-controller/internal/core/entity"

type RepositoryManager interface {
	BeginTran() (RepositoryManager, error)
	CommitTran() error
	RollbackTran() error
	Transaction(callback func(RepositoryManager) error) error

	GetUserRepository() UserRepository
}

type UserRepository interface {
	Get(Id int) (entity.User, error)
	GetByUsername(username string, sourceId int) (entity.User, error)
	GetList(filters entity.UserFilters) ([]entity.User, int, error)
	Save(*entity.User) error
	Delete(id int) error
}
