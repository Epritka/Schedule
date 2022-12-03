package user

import "engine/internal/core/interfaces"

type handlers struct {
	userUseCase interfaces.UserUseCase
}

func NewHandlers(
	userUseCase interfaces.UserUseCase,
) handlers {
	return handlers{
		userUseCase: userUseCase,
	}
}
