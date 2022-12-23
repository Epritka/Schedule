package user

import (
	"strconv"

	"user-manager/infrastructure/serializers"

	"github.com/gin-gonic/gin"
)

func (handlers *handlers) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	user, err := handlers.userUseCase.Get(id)
	if err != nil {
		c.JSON(serializers.NotFoundHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(user))
}
