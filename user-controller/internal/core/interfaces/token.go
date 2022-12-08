package interfaces

import (
	"time"

	"user-controller/internal/core/entity"
)

type TokenGenerator interface {
	NewAccessToken(user entity.User, ttl time.Duration) (string, error)
	NewRefreshToken(userId int, isRemember bool, ttl time.Duration) (string, error)
	GetJwks() entity.Jwks
	VerifyAccessToken(token string) (entity.User, error)
	VerifyRefreshToken(token string) (int, bool, error)
}
