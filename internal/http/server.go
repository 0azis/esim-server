package http

import (
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

func New(addr string) Server {
	return &server{
		e:    echo.New(),
		addr: addr,
	}
}

func (s server) Run() error {
	return s.e.Start(s.addr)
}

func (s server) ApiRouter() *echo.Group {
	return s.e.Group(fmt.Sprintf("/api/%s", apiVersion))
}
