package http

import (
	"fmt"

	"engine/internal/core/interfaces"

	"github.com/gin-gonic/gin"
)

type httpServer struct {
	port       int
	router     *gin.Engine
	dayUseCase interfaces.DayUseCase
}

func New(
	isDebug bool,
	port int,
	dayUseCase interfaces.DayUseCase,
) *httpServer {
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	// TODO move to Config
	if port == 0 {
		port = 2513
	}

	s := &httpServer{
		router:     gin.Default(),
		dayUseCase: dayUseCase,
		port:       port,
	}
	s.SetRoutes()
	return s
}

func (s *httpServer) Run() {
	s.router.Run(fmt.Sprintf(":%d", s.port))
}
