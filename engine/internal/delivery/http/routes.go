package http

import (
	"engine/internal/delivery/http/internal/user"
)

func (s *httpServer) SetRoutes() {
	userHandlers := user.New(s.userUseCase)

	apiV1 := s.router.Group("/api/v1/")
	{

		user := apiV1.Group("/days/")
		{
			user.GET("/", userHandlers.GetList)
			user.POST("/", userHandlers.Create)
			id := user.Group("/:id/")
			{
				id.GET("/", userHandlers.Get)
				id.PATCH("/", userHandlers.Update)
				id.DELETE("/", userHandlers.Delete)
			}
		}
	}
}
