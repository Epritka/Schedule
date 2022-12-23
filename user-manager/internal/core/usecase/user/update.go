package user

import (
	"user-manager/internal/core/entity"
)

func (usecase *useCase) Update(id, tgUserId int) (entity.User, error) {
	userRepository := usecase.repositoryManager.GetUserRepository()
	user, err := userRepository.Get(id)

	if err != nil {
		return entity.User{}, entity.NotFoundError
	}

	err = userRepository.Save(&user)
	if err != nil {
		return entity.User{}, entity.DatabaseError
	}

	return user, nil
}
