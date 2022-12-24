package group

import (
	"engine/infrastructure/serializers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (handlers *handlers) GetByName(c *gin.Context) {
	name := c.Param("name")

	fmt.Println(name)
	group, err := handlers.groupUseCase.GetByName(name)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	if group == nil {
		c.JSON(serializers.NotFoundHttpResponce(fmt.Errorf("group not found"), nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(group))
}
