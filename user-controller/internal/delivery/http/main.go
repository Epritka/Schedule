package http

import (
	"fmt"

	"user-controller/internal/core/interfaces"

	"github.com/gin-gonic/gin"
)

type httpServer struct {
	port        int
	router      *gin.Engine
	userUseCase interfaces.UserUseCase
}

func New(
	isProduction bool,
	port int,
	userUseCase interfaces.UserUseCase,
) *httpServer {
	if isProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	if port == 0 {
		port = 8080
	}

	s := &httpServer{
		router:      gin.Default(),
		userUseCase: userUseCase,
		port:        port,
	}
	s.SetRoutes()
	return s
}

func (s *httpServer) Run() {
	s.router.Run(fmt.Sprintf(":%d", s.port))
}
