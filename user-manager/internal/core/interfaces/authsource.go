package interfaces

import "user-manager/internal/core/entity"

type AuthSource interface {
	GetUser(username, password string) (entity.User, error)
}

type AuthSources interface {
	Get(source entity.AuthSource) (AuthSource, error)
	Validate(source entity.AuthSource) bool
}
