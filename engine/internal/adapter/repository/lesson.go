package repository

import (
	"engine/internal/adapter/repository/model"
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"
	"fmt"

	"github.com/go-pg/pg/orm"
)

type lessonRepository struct {
	DB orm.DB
}

func NewLessonRepository(db orm.DB) interfaces.LessonRepository {
	return &lessonRepository{
		DB: db,
	}
}

func (r *lessonRepository) Save(lesson *entity.Lesson) (*entity.Lesson, error) {
	fmt.Println(*lesson)
	model := model.NewLesson(*lesson)
	exist, err := r.DB.Model(&model).WherePK().Exists()
	if err != nil {
		return lesson, err
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
		return lesson, err
	}

	lesson.Id = model.Id

	return lesson, nil

}
