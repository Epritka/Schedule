package user

import (
	"user-manager/internal/core/entity"
)

func (usecase *useCase) GetList(filters entity.UserFilters) ([]entity.User, int, error) {
	userRepository := usecase.repositoryManager.GetUserRepository()

	if filters.Limit > 50 {
		filters.Limit = 50
	}

	userList, count, err := userRepository.GetList(filters)
	if err != nil {
		return []entity.User{}, 0, entity.DatabaseError
	}
	return userList, count, nil
}
