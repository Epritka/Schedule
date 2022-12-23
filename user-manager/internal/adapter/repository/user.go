package repository

import (
	"user-manager/internal/adapter/repository/model"
	"user-manager/internal/core/entity"
	"user-manager/internal/core/interfaces"

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
	user := model.User{
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

func (r *userRepository) GetByTelegramUserId(tgId int) (*entity.User, error) {
	user := model.User{}

	query := r.DB.Model(&user).
		Relation("Roles").
		Relation("AuthSource").
		Where("telegram_user_id = ?", tgId)

	err := query.
		Order("id ASC").
		Select()

	if err != nil {
		return nil, err
	}

	userEntity := user.Entity()

	return &userEntity, nil
}

func (r *userRepository) GetList() ([]entity.User, int, error) {
	list := model.UserList{}
	query := r.DB.Model(&list)

	count, err := query.Count()
	if err != nil {
		return list.Entity(), 0, err
	}

	err = query.
		Order("id ASC").
		Select()

	if err != nil {
		return list.Entity(), 0, err
	}

	return list.Entity(), count, nil
}

func (r *userRepository) Save(user *entity.User) error {
	model := model.NewUser(*user)
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
	_, err := r.DB.Model(&model.User{Id: id}).
		WherePK().
		Delete()
	if err != nil {
		return err
	}
	return nil
}
