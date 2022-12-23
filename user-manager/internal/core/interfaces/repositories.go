package interfaces

import "user-manager/internal/core/entity"

type RepositoryManager interface {
	BeginTran() (RepositoryManager, error)
	CommitTran() error
	RollbackTran() error
	Transaction(callback func(RepositoryManager) error) error

	GetUserRepository() UserRepository
}

type UserRepository interface {
	Get(Id int) (entity.User, error)
	GetByTelegramUserId(tgId int) (*entity.User, error)
	GetList() ([]entity.User, int, error)
	Save(*entity.User) error
	Delete(id int) error
}
