package schedule

import (
	"engine/infrastructure/serializers"
	"engine/internal/core/entity"

	"github.com/gin-gonic/gin"
)

func (handlers *handlers) Save(c *gin.Context) {
	var request []entity.Schedule
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	err := handlers.scheduleUseCase.Save(request)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(nil))
}
