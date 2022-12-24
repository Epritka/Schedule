package repository

import (
	"engine/internal/adapter/repository/model"
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"

	"github.com/go-pg/pg/orm"
)

type educationalRepository struct {
	DB orm.DB
}

func NewEducationalRepository(db orm.DB) interfaces.EducationalRepository {
	return &educationalRepository{
		DB: db,
	}
}

func (r *educationalRepository) SaveEducationalInstitution(name string) (entity.EducationalInstitution, error) {
	educationalInstitution := entity.EducationalInstitution{Name: name}

	model := model.NewEducationalInstitution(educationalInstitution)
	exist, err := r.DB.Model(&model).WherePK().Exists()
	if err != nil {
		return educationalInstitution, err
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
		return educationalInstitution, err
	}

	educationalInstitution.Id = model.Id

	return educationalInstitution, nil
}

func (r *educationalRepository) SaveFaculty(name string) (entity.Faculty, error) {
	faculty := entity.Faculty{Name: name}

	model := model.NewFaculty(faculty)
	exist, err := r.DB.Model(&model).WherePK().Exists()
	if err != nil {
		return faculty, err
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
		return faculty, err
	}

	faculty.Id = model.Id

	return faculty, nil
}

func (r *educationalRepository) SaveYear(name string) (entity.Year, error) {
	year := entity.Year{Name: name}

	model := model.NewUYear(year)
	exist, err := r.DB.Model(&model).WherePK().Exists()
	if err != nil {
		return year, err
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
		return year, err
	}

	year.Id = model.Id

	return year, nil
}
