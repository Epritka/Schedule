package user

import (
	"errors"
	"strconv"

	"user-manager/infrastructure/serializers"
	"user-manager/infrastructure/validator"

	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	TelegramUserId int `json:"telegramUserId" validate:"required"`
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
		request.TelegramUserId,
	)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(user))
}
