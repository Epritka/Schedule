package repository

import (
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"
)

type userRepository struct {
}

func NewUserRepository() interfaces.UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetByName(string) (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) Update(entity.User) error {
	return nil
}
