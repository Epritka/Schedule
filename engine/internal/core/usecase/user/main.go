package user

import "engine/internal/core/interfaces"

type useCase struct {
	repository interfaces.RepositoryManager
}

func NewUserUseCase(repository interfaces.RepositoryManager) interfaces.UserUseCase {
	return &useCase{repository: repository}
}
