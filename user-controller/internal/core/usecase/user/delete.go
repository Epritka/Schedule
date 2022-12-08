package user

import (
	"user-controller/internal/core/entity"
)

func (usecase *useCase) Delete(id int) error {
	userRepository := usecase.repositoryManager.GetUserRepository()
	err := userRepository.Delete(id)
	if err != nil {
		return entity.DatabaseError
	}
	return nil
}
