package user

import "engine/internal/core/interfaces"

type UserUseCase struct {
	repository interfaces.RepositoryManager
}

func NewUserUseCase(repository interfaces.RepositoryManager) interfaces.UserUseCase {
	return &UserUseCase{repository: repository}
}
