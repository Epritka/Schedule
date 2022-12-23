package user

import (
	"errors"

	"user-manager/infrastructure/serializers"
	"user-manager/infrastructure/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Email           string `json:"email" validate:"required,email,max=100"`
	Password        string `json:"password" validate:"required,max=100"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,max=100,eqfield=Password"`
	FirstName       string `json:"firstName" validate:"required,max=100"`
	LastName        string `json:"lastName" validate:"required,max=100"`
}

func (handlers *handlers) Create(c *gin.Context) {
	var request createRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	validator := validator.New()
	if err := validator.Struct(request); err != nil {
		c.JSON(serializers.BadRequestHttpResponce(
			errors.New("validation errors"),
			validator.GetErrors(err),
		))
		return
	}

	user, err := handlers.userUseCase.Create(
		request.Email,
		request.Password,
		request.FirstName,
		request.LastName,
	)
	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}
	c.JSON(serializers.SuccessHttpResponce(user))
}
