package student

import (
	"engine/infrastructure/serializers"
	"engine/internal/core/entity"

	"github.com/gin-gonic/gin"
)

func (handlers *handlers) Save(c *gin.Context) {
	var request entity.Student
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	student, err := handlers.studentUseCase.Save(&request)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(student))
}
