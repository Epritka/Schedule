package repository

import (
	"engine/internal/adapter/repository/model"
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"

	"github.com/go-pg/pg/orm"
)

type dayRepository struct {
	DB orm.DB
}

func NewDayRepository(db orm.DB) interfaces.DayRepository {
	return &dayRepository{
		DB: db,
	}
}

func (r *dayRepository) GetLessons(
	weekType string,
	weekDay entity.Weekday,
	groupId int,
) (*entity.Day, error) {

	day := model.Day{}

	query := r.DB.Model(&day).
		Relation("Lessons").
		Where("group_id = ?", groupId).
		Where("week_type = ?", weekType).
		Where("number = ?", weekDay)

	err := query.Select()

	if err != nil {
		if err != nil {
			if err.Error() == "pg: no rows in result set" {
				return nil, nil
			}

			return nil, err
		}
	}

	d := day.Entity()
	return &d, nil
}

func (r *dayRepository) Save(weekType string, number, groupId int) (*entity.Day, error) {
	day := entity.Day{Number: number, WeekType: weekType, GroupId: groupId}
	model := model.Day{Number: number, WeekType: weekType, GroupId: groupId}

	exist, err := r.DB.Model(&model).WherePK().Exists()
	if err != nil {
		return &day, err
	}

	if exist {
		_, err = r.DB.Model(&model).
			WherePK().
			Update()
	} else {
		_, err = r.DB.Model(&model).
			Insert()
	}

	if err != nil {
		return &day, err
	}

	day.Id = model.Id

	return &day, nil

}
