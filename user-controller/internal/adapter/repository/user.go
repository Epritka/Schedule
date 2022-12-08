package repository

import (
	"user-controller/internal/adapter/repository/models"
	"user-controller/internal/core/entity"
	"user-controller/internal/core/interfaces"

	"github.com/go-pg/pg/orm"
)

type userRepository struct {
	DB orm.DB
}

func NewUserRepository(
	DB orm.DB,
) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r *userRepository) Get(id int) (entity.User, error) {
	user := models.User{
		Id: id,
	}
	err := r.DB.Model(&user).
		Relation("Roles").
		Relation("AuthSource").
		WherePK().
		Select()
	if err != nil {
		return user.Entity(), err
	}
	return user.Entity(), nil
}

func (r *userRepository) GetByUsername(username string, sourceId int) (entity.User, error) {
	user := models.User{}

	query := r.DB.Model(&user).
		Relation("Roles").
		Relation("AuthSource").
		Where("email = ?", username)

	if sourceId != 0 {
		query = query.Where("auth_source_id = ?", sourceId)
	}

	err := query.
		Order("id ASC").
		Select()
	if err != nil {
		return user.Entity(), err
	}
	return user.Entity(), nil
}

func (r *userRepository) GetList(filters entity.UserFilters) ([]entity.User, int, error) {
	list := models.UserList{}
	query := r.DB.Model(&list).
		Relation("AuthSource").
		Relation("Roles")

	if len(filters.Ids) > 0 {
		query.WhereIn("id in (?)", filters.Ids)
	}

	if filters.Email != "" {
		query.Where("email like ?", "%"+filters.Email+"%")
	}

	if filters.FullName != "" {
		query.Where("concat_ws(' ', first_name, last_name) like ?", "%"+filters.FullName+"%")
	}

	count, err := query.Count()
	if err != nil {
		return list.Entity(), 0, err
	}

	err = query.
		Limit(filters.Limit).
		Offset(filters.Offset).
		Order("id ASC").
		Select()
	if err != nil {
		return list.Entity(), 0, err
	}
	return list.Entity(), count, nil
}

func (r *userRepository) Save(user *entity.User) error {
	model := models.NewUser(*user)
	exist, err := r.DB.Model(&model).WherePK().Exists()
	if err != nil {
		return err
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
		return err
	}

	user.Id = model.Id
	return nil
}

func (r *userRepository) Delete(id int) error {
	_, err := r.DB.Model(&models.User{Id: id}).
		WherePK().
		Delete()
	if err != nil {
		return err
	}
	return nil
}
