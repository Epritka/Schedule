package models

import (
	"user-controller/infrastructure/convert"
	"user-controller/internal/core/entity"
)

type UserList []User

type User struct {
	tableName struct{} `sql:"app_user"`

	Id           int
	AuthSourceId int
	Email        string
	FirstName    string
	LastName     string
	Password     string
	IsSuperuser  bool                `pg:",use_zero" sql:",notnull"`
	IsActive     bool                `pg:",use_zero" sql:",notnull"`
	Roles        RoleList            `pg:",many2many:user_roles"`
	Permissions  []entity.Permission `pg:",array"`
	AuthSource   AuthSource          `pg:"rel:has-one"`
}

func NewUser(value entity.User) User {
	return convert.Convert[entity.User, User](value)
}

func (value *User) Entity() entity.User {
	result := convert.DeConvert[User, entity.User](*value)
	result.Roles = value.Roles.Entity()
	result.AuthSource = value.AuthSource.Entity()
	return result
}

func (list *UserList) Entity() []entity.User {
	result := []entity.User{}
	for _, value := range *list {
		result = append(result, value.Entity())
	}
	return result
}
