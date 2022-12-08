package models

import (
	"time"

	"user-controller/infrastructure/convert"
	"user-controller/internal/core/entity"
)

type ApiTokenList []ApiToken

type ApiToken struct {
	Id          int
	UserId      int
	Token       string
	Description string
	Ttl         *time.Time
}

func NewApiToken(value entity.ApiToken) ApiToken {
	return convert.Convert[entity.ApiToken, ApiToken](value)
}

func (value *ApiToken) Entity() entity.ApiToken {
	result := convert.DeConvert[ApiToken, entity.ApiToken](*value)
	return result
}

func (list *ApiTokenList) Entity() []entity.ApiToken {
	result := []entity.ApiToken{}
	for _, value := range *list {
		result = append(result, value.Entity())
	}
	return result
}
