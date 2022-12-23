package model

import (
	"user-manager/infrastructure/convert"
	"user-manager/internal/core/entity"
)

type UserList []User

type User struct {
	tableName struct{} `sql:"app_user"`

	Id             int
	TelegramUserId int
	Email          string
	FirstName      string
	LastName       string
	Password       string
	IsSuperuser    bool `pg:",use_zero" sql:",notnull"`
	IsActive       bool `pg:",use_zero" sql:",notnull"`
}

func NewUser(value entity.User) User {
	return convert.Convert[entity.User, User](value)
}

func (value *User) Entity() entity.User {
	result := convert.DeConvert[User, entity.User](*value)
	return result
}

func (list *UserList) Entity() []entity.User {
	result := []entity.User{}
	for _, value := range *list {
		result = append(result, value.Entity())
	}
	return result
}
