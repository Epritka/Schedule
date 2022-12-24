package group

import (
	"engine/internal/core/interfaces"
)

type handlers struct {
	groupUseCase interfaces.GroupUseCase
}

func New(groupUseCase interfaces.GroupUseCase) *handlers {
	return &handlers{
		groupUseCase: groupUseCase,
	}
}
