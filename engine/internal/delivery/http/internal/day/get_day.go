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

func (handlers *handlers) GetList(c *gin.Context) {
	var filters DayFilters

	c.BindQuery(&filters)

	users, count, err := handlers.userUseCase.GetList(entity.UserFilters{
		Ids:      filters.Ids,
		Email:    filters.Email,
		Limit:    filters.Limit,
		Offset:   filters.Offset,
		FullName: filters.FullName,
	})
	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
	}

	code, data := serializers.SuccessHttpResponce(users)
	data["count"] = count
	c.JSON(code, data)
}
