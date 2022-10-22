package schedule

import "engine/internal/core/entity"

func (usecase *ScheduleUseCase) Create(schedules []entity.Schedule) error {
	scheduleRepo := usecase.repository.GetScheduleRepository()

	err := scheduleRepo.Create(schedules)
	return err
}
