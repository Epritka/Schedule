package day

import (
	"engine/infrastructure/serializers"
	"engine/internal/core/entity"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DayFilters struct {
	GroupId int    `form:"groupId"`
	Date    string `form:"date"`
}

func (handlers *handlers) Get(c *gin.Context) {
	var filters DayFilters
	c.ShouldBindQuery(&filters)

	day, err := handlers.dayUseCase.Get(
		entity.DayFilter{
			GroupId: filters.GroupId,
			Date:    filters.Date,
		},
	)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	if day == nil {
		c.JSON(serializers.NotFoundHttpResponce(fmt.Errorf("lessons not found"), nil))
		return
	}

	c.JSON(serializers.SuccessHttpResponce(day))
}
