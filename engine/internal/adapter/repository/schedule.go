package repository

import (
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"
)

type scheduleRepository struct {
}

func NewScheduleRepository() interfaces.ScheduleRepository {
	return &scheduleRepository{}
}

func (r *scheduleRepository) Create(schedules []entity.Schedule) error {
	return nil
}

func (r *scheduleRepository) GetGroupId(groupName string) *int {
	return nil
}
