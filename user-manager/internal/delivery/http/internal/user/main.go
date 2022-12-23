package user

import (
	"user-manager/internal/core/interfaces"
)

type handlers struct {
	userUseCase interfaces.UserUseCase
}

func New(userUseCase interfaces.UserUseCase) *handlers {
	return &handlers{
		userUseCase: userUseCase,
	}
}
