package interfaces

import (
	"user-controller/internal/core/entity"
)

type UserUseCase interface {
	GetList(filters entity.UserFilters) ([]entity.User, int, error)
	Get(id int) (entity.User, error)
	Create(email, password, firstName, lastName string) (entity.User, error)
	Update(id int, email, firstName, lastName string) (entity.User, error)
	Delete(id int) error
}
