package user

import (
	interfaces "user-controller/internal/core/interfaces"
)

type useCase struct {
	repositoryManager interfaces.RepositoryManager
	cryptographer     interfaces.Cryptographer
	logger            interfaces.Logger
}

func New(
	repositoryManager interfaces.RepositoryManager,
	cryptographer interfaces.Cryptographer,
	logger interfaces.Logger,
) interfaces.UserUseCase {
	return &useCase{
		repositoryManager: repositoryManager,
		cryptographer:     cryptographer,
		logger:            logger,
	}
}
