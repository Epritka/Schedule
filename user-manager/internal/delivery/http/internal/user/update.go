package user

import (
	"errors"
	"strconv"

	"user-manager/infrastructure/serializers"
	"user-manager/infrastructure/validator"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	Email     string `json:"email" validate:"required,email,max=100"`
	FirstName string `json:"firstName" validate:"required,max=100"`
	LastName  string `json:"lastName" validate:"required,max=100"`
}

func (handlers *handlers) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	var request updateRequest
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

	user, err := handlers.userUseCase.Update(
		id,
		request.Email,
		request.FirstName,
		request.LastName,
	)
	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}
	c.JSON(serializers.SuccessHttpResponce(user))
}
