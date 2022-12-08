package models

import (
	"user-controller/infrastructure/convert"
	"user-controller/internal/core/entity"
)

type RoleList []Role

type Role struct {
	Id          int
	Name        string
	Permissions []entity.Permission `pg:",array"`
}

func NewRole(value entity.Role) Role {
	return convert.Convert[entity.Role, Role](value)
}

func (value *Role) Entity() entity.Role {
	return convert.DeConvert[Role, entity.Role](*value)
}

func (list *RoleList) Entity() []entity.Role {
	return convert.DeConvertList[Role, entity.Role](*list)
}
