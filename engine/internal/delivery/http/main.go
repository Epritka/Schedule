package http

import (
	"fmt"

	"engine/internal/core/interfaces"

	"github.com/gin-gonic/gin"
)

type httpServer struct {
	port            int
	router          *gin.Engine
	dayUseCase      interfaces.DayUseCase
	scheduleUseCase interfaces.ScheduleUseCase
	studentUseCase  interfaces.StudentUseCase
	groupUseCase    interfaces.GroupUseCase
}

func New(
	isDebug bool,
	port int,
	dayUseCase interfaces.DayUseCase,
	scheduleUseCase interfaces.ScheduleUseCase,
	studentUseCase interfaces.StudentUseCase,
	groupUseCase interfaces.GroupUseCase,
) *httpServer {
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &httpServer{
		router:          gin.Default(),
		dayUseCase:      dayUseCase,
		scheduleUseCase: scheduleUseCase,
		studentUseCase:  studentUseCase,
		groupUseCase:    groupUseCase,
		port:            port,
	}
	s.SetRoutes()
	return s
}

func (s *httpServer) Run() {
	s.router.Run(fmt.Sprintf(":%d", s.port))
}
