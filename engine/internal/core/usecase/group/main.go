package group

import "engine/internal/core/interfaces"

type useCase struct {
	repository interfaces.RepositoryManager
	logger     interfaces.Logger
}

func New(repository interfaces.RepositoryManager, logger interfaces.Logger) interfaces.GroupUseCase {
	return &useCase{
		repository: repository,
		logger:     logger,
	}
}
