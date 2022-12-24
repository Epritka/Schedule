package repository

import (
	"engine/internal/adapter/repository/model"
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"

	"github.com/go-pg/pg/orm"
)

type studentRepository struct {
	DB orm.DB
}

func NewStudentRepository(db orm.DB) interfaces.StudentRepository {
	return &studentRepository{
		DB: db,
	}
}

func (r *studentRepository) Save(student *entity.Student) (*entity.Student, error) {
	model := model.NewStudent(*student)
	exist, err := r.DB.Model(&model).WherePK().Exists()
	if err != nil {
		return student, err
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
		return student, err
	}

	student.Id = model.Id

	return student, nil
}

func (r *studentRepository) Get(id int) (*entity.Student, error) {
	model := model.Student{
		Id: id,
	}

	err := r.DB.Model(&model).
		WherePK().
		Select()

	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, nil
		}

		return nil, err
	}

	student := model.Entity()
	return &student, nil
}

func (r *studentRepository) GetList() ([]entity.Student, error) {
	list := model.StudentList{}
	query := r.DB.Model(&list).
		Order("id DESC")

	err := query.Select()

	if err != nil {
		return list.Entity(), err
	}

	return list.Entity(), nil
}
