package day

import (
	"engine/internal/core/interfaces"
)

type handlers struct {
	dayUseCase interfaces.DayUseCase
}

func New(dayUseCase interfaces.DayUseCase) *handlers {
	return &handlers{
		dayUseCase: dayUseCase,
	}
}
