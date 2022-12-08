package models

import (
	"user-controller/infrastructure/convert"
	"user-controller/internal/core/entity"
)

type AuthSourceList []AuthSource

type AuthSource struct {
	Id         int                   `json:"id"`
	Name       string                `json:"name"`
	SourceType entity.AuthSourceType `json:"sourceType"`
	Url        string                `json:"url"`
	IsActive   bool                  `json:"isActive"`
	Meta       map[string]any        `json:"meta"`
}

func NewAuthSource(value entity.AuthSource) AuthSource {
	return convert.Convert[entity.AuthSource, AuthSource](value)
}

func (value *AuthSource) Entity() entity.AuthSource {
	result := convert.DeConvert[AuthSource, entity.AuthSource](*value)
	return result
}

func (list *AuthSourceList) Entity() []entity.AuthSource {
	result := []entity.AuthSource{}
	for _, value := range *list {
		result = append(result, value.Entity())
	}
	return result
}
