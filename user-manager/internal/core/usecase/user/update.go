package user

import (
	"user-manager/internal/core/entity"
)

func (usecase *useCase) Update(id int, email, firstName, lastName string) (entity.User, error) {
	userRepository := usecase.repositoryManager.GetUserRepository()
	user, err := userRepository.Get(id)
	if err != nil {
		return entity.User{}, entity.NotFoundError
	}

	if email != "" {
		user.Email = email
	}

	if firstName != "" {
		user.FirstName = firstName
	}

	if lastName != "" {
		user.LastName = lastName
	}

	err = userRepository.Save(&user)
	if err != nil {
		return entity.User{}, entity.DatabaseError
	}

	return user, nil
}
