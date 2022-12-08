package user

import (
	"user-controller/internal/core/entity"
)

func (usecase *useCase) Create(email, password, firstName, lastName string) (entity.User, error) {
	userRepository := usecase.repositoryManager.GetUserRepository()
	eUser, _ := userRepository.GetByUsername(email, 0)
	if eUser.Id != 0 {
		return entity.User{}, entity.UserAlreadyExistError
	}

	user := entity.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	if !user.ValidPassword(password) {
		return entity.User{}, entity.InvalidAuthDataError
	}

	cryptPassword, err := usecase.cryptographer.Encrypt(password)
	if err != nil {
		return entity.User{}, entity.InvalidAuthDataError
	}

	user.Password = cryptPassword
	err = userRepository.Save(&user)
	if err != nil {
		return entity.User{}, entity.DatabaseError
	}

	return user, nil
}
