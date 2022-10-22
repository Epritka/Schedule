package telegram

import (
	"engine/infrastructure/config"
)

type HttpServer struct {
	Url string
}

func NewHttpServer(config config.Config) (*HttpServer, error) {
	return nil, nil
}
