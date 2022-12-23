package user

import (
	"user-manager/infrastructure/serializers"
	"user-manager/internal/core/entity"

	"github.com/gin-gonic/gin"
)

type UserFilters struct {
	Ids      []int  `form:"ids"`
	Email    string `form:"email"`
	FullName string `form:"fullName"`

	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

func (handlers *handlers) GetList(c *gin.Context) {
	var filters UserFilters

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
