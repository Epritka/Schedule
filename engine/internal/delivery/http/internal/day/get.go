package day

import (
	"engine/infrastructure/serializers"
	"engine/internal/core/entity"

	"github.com/gin-gonic/gin"
)

type DayFilters struct {
	GroupId int    `json:"groupId"`
	Date    string `json:"date"`
}

func (handlers *handlers) Get(c *gin.Context) {
	var filters DayFilters

	c.BindQuery(&filters)

	day, err := handlers.dayUseCase.Get(
		entity.DayFilter{
			GroupId: filters.GroupId,
			Date:    filters.Date,
		},
	)

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
	}

	code, data := serializers.SuccessHttpResponce(day)
	c.JSON(code, data)
}
