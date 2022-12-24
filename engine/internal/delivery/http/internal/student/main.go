package student

import (
	"engine/internal/core/interfaces"
)

type handlers struct {
	studentUseCase interfaces.StudentUseCase
}

func New(studentUseCase interfaces.StudentUseCase) *handlers {
	return &handlers{
		studentUseCase: studentUseCase,
	}
}
