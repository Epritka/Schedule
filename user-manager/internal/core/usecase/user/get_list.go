package user

import (
	"fmt"
	"user-manager/internal/core/entity"
)

func (usecase *useCase) GetList() ([]entity.User, int, error) {
	userRepository := usecase.repositoryManager.GetUserRepository()

	userList, count, err := userRepository.GetList()
	if err != nil {
		fmt.Println(err)
		return []entity.User{}, 0, entity.DatabaseError
	}

	return userList, count, nil
}
