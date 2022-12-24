package repository

import (
	"engine/internal/adapter/repository/model"
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"

	"github.com/go-pg/pg/orm"
)

type groupRepository struct {
	DB orm.DB
}

func NewGroupRepository(db orm.DB) interfaces.GroupRepository {
	return &groupRepository{
		DB: db,
	}
}

func (r *groupRepository) Save(group *entity.Group) (*entity.Group, error) {
	model := model.NewGroup(*group)
	exist, err := r.DB.Model(&model).WherePK().Exists()
	if err != nil {
		return group, err
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
		return group, err
	}

	group.Id = model.Id

	return group, nil

}

func (r *groupRepository) Get(id int) (*entity.Group, error) {
	model := model.Group{
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

	group := model.Entity()
	return &group, nil
}

func (r *groupRepository) GetByName(name string) (*entity.Group, error) {
	model := model.Group{}

	err := r.DB.Model(&model).
		Where("name like ?", name).
		Select()

	if err != nil {
		if err.Error() == "pg: no rows in result set" {
			return nil, nil
		}

		return nil, err
	}

	group := model.Entity()
	return &group, nil
}
