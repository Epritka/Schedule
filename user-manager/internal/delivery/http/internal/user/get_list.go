package user

import (
	"user-manager/infrastructure/serializers"

	"github.com/gin-gonic/gin"
)

func (handlers *handlers) GetList(c *gin.Context) {

	users, count, err := handlers.userUseCase.GetList()

	if err != nil {
		c.JSON(serializers.BadRequestHttpResponce(err, nil))
		return
	}

	code, data := serializers.SuccessHttpResponce(users)
	data["count"] = count
	c.JSON(code, data)
}
