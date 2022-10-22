package user

import (
	"engine/internal/core/entity"
	"engine/internal/core/interfaces"
)

type UserRepository struct {
	users []entity.User
}

func NewUserRepository() interfaces.UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Update(login string, groupId int) error {
	for _, u := range r.users {
		if u.Login == login {
			u.GroupId = groupId
			return nil
		}
	}

	r.users = append(r.users, entity.User{
		Login:   login,
		GroupId: groupId,
	})

	return nil
}
