package http

import (
	"engine/internal/delivery/http/internal/day"
	"engine/internal/delivery/http/internal/group"
	"engine/internal/delivery/http/internal/schedule"
	"engine/internal/delivery/http/internal/student"
)

func (s *httpServer) SetRoutes() {
	dayHandlers := day.New(s.dayUseCase)
	scheduleHandlers := schedule.New(s.scheduleUseCase)
	studentHandlers := student.New(s.studentUseCase)
	groupHandlers := group.New(s.groupUseCase)

	apiV1 := s.router.Group("/api/v1/")
	{
		day := apiV1.Group("/day/")
		{
			day.GET("/", dayHandlers.Get)
		}

		schedule := apiV1.Group("/schedule/")
		{
			schedule.POST("/", scheduleHandlers.Save)
		}

		group := apiV1.Group("/group/")
		{
			group.GET("/", groupHandlers.GetByName)
		}

		student := apiV1.Group("/student/")
		{
			student.POST("/", studentHandlers.Save)
			student.PATCH("/", studentHandlers.Save)
			student.GET("/", studentHandlers.GetList)
			id := student.Group("/:id/")
			{
				id.GET("/", studentHandlers.Get)
			}
		}
	}
}
