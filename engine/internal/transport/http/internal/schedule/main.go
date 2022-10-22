package schedule

import "engine/internal/core/interfaces"

type handlers struct {
	scheduleUseCase interfaces.ScheduleUseCase
}

func NewHandlers(
	scheduleUseCase interfaces.ScheduleUseCase,
) handlers {
	return handlers{
		scheduleUseCase: scheduleUseCase,
	}
}
