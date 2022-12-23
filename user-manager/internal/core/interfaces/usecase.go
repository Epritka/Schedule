package interfaces

import (
	"user-manager/internal/core/entity"
)

type UserUseCase interface {
	GetList() ([]entity.User, int, error)
	Get(id int) (entity.User, error)
	Create(telegramUserId int) (entity.User, error)
	Update(id, telegramUserId int) (entity.User, error)
	Delete(id int) error
}
