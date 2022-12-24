package group

import (
	"engine/infrastructure/serializers"

	"github.com/gin-gonic/gin"
)

func (handlers *handlers) GetByName(c *gin.Context) {
	name := c.Param("name")

	group, err := handlers.groupUseCase.GetByName(name)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(group))
}
