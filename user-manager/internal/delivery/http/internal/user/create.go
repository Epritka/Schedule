package user

import (
	"errors"

	"user-manager/infrastructure/serializers"
	"user-manager/infrastructure/validator"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	TelegramUserId int `json:"telegramUserId" validate:"required"`
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
		request.TelegramUserId,
	)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(user))
}
