package user

import (
	"user-manager/internal/core/entity"
)

func (usecase *useCase) Create(tgUserId int) (entity.User, error) {
	userRepository := usecase.repositoryManager.GetUserRepository()
	eUser, err := userRepository.GetByTelegramUserId(tgUserId)
	if err != nil {
		usecase.logger.Error(err.Error())
		return entity.User{}, entity.DatabaseError
	}

	if eUser != nil {
		return entity.User{}, entity.UserAlreadyExistError
	}

	user := entity.User{
		TelegramUserId: tgUserId,
	}

	err = userRepository.Save(&user)
	if err != nil {
		usecase.logger.Error(err.Error())
		return entity.User{}, entity.DatabaseError
	}

	return user, nil
}
