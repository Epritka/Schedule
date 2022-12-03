package http

import (
	"engine/infrastructure/config"
	"engine/internal/core/interfaces"
	"engine/internal/delivery/http/internal/schedule"
	"engine/internal/delivery/http/internal/user"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	port   int
	router *mux.Router
}

func NewHttpServer(
	config config.Config,
	scheduleUseCase interfaces.ScheduleUseCase,
	userUseCase interfaces.UserUseCase,
) (*HttpServer, error) {
	router := mux.NewRouter()

	userHandlers := user.NewHandlers(userUseCase)

	router.HandleFunc("/user/login", userHandlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/user/{login}", userHandlers.Login).Methods(http.MethodPost)

	scheduleHandlers := schedule.NewHandlers(scheduleUseCase)

	router.HandleFunc("/schedule/day", scheduleHandlers.GetScheduleForDay).Methods(http.MethodGet)
	router.HandleFunc("/schedule/upload", scheduleHandlers.UploadNewSchedule).Methods(http.MethodPost)

	if config.IsDebug {
		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			tpl, _ := route.GetPathTemplate()
			met, _ := route.GetMethods()
			fmt.Println(tpl, met)

			return nil
		})
	}

	return &HttpServer{port: config.Port, router: router}, nil
}

func (s *HttpServer) Start() error {
	return http.ListenAndServe(":2513", s.router)
}
