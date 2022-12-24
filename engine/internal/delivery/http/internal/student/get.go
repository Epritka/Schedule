package student

import (
	"engine/infrastructure/serializers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DayFilters struct {
	GroupId int    `json:"groupId"`
	Date    string `json:"date"`
}

func (handlers *handlers) GetList(c *gin.Context) {
	students, err := handlers.studentUseCase.GetList()

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(students))
}

func (handlers *handlers) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	students, err := handlers.studentUseCase.Get(id)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(students))
}
