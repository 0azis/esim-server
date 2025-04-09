package http

import (
	"esim/config"
	"fmt"

	"github.com/labstack/echo/v4"
)

const apiVersion = "v1"

type Server interface {
	Run() error
	ApiRouter() *echo.Group
}

type server struct {
	e *echo.Echo

	addr string
}

func New(cfg config.Config) Server {
	return &server{
		e:    echo.New(),
		addr: cfg.Http.Addr(),
	}
}

func (s server) Run() error {
	return s.e.Start(s.addr)
}

func (s server) ApiRouter() *echo.Group {
	return s.e.Group(fmt.Sprintf("/api/%s", apiVersion))
}
