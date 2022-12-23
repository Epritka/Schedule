package user

import (
	"strconv"

	"user-manager/infrastructure/serializers"

	"github.com/gin-gonic/gin"
)

func (handlers *handlers) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	err = handlers.userUseCase.Delete(id)
	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(nil))
}
