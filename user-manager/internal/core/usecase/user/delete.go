package user

import (
	"user-manager/internal/core/entity"
)

func (usecase *useCase) Delete(id int) error {
	userRepository := usecase.repositoryManager.GetUserRepository()
	err := userRepository.Delete(id)
	if err != nil {
		return entity.DatabaseError
	}
	return nil
}
