package user

import (
	interfaces "user-manager/internal/core/interfaces"
)

type useCase struct {
	repositoryManager interfaces.RepositoryManager
	logger            interfaces.Logger
}

func New(
	repositoryManager interfaces.RepositoryManager,
	logger interfaces.Logger,
) interfaces.UserUseCase {
	return &useCase{
		repositoryManager: repositoryManager,
		logger:            logger,
	}
}
