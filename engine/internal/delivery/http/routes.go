package http

import (
	"engine/internal/delivery/http/internal/day"
)

func (s *httpServer) SetRoutes() {
	dayHandlers := day.New(s.dayUseCase)

	apiV1 := s.router.Group("/api/v1/")
	{
		schedule := apiV1.Group("/schedule/")
		{
			day := schedule.Group("/day/")
			{
				day.GET("/", dayHandlers.Get)
			}
		}
	}
}
