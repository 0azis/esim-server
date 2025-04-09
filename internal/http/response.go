package http

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Data any  `json:"data"`
	Meta meta `json:"meta"`
}

type meta struct {
	Descipriton string `json:"description"`
	Version     string `json:"version"`
	Timestamp   string `json:"timestamp"`
}

func NewResponse(c echo.Context, statusCode int, data any) error {
	return c.JSON(statusCode, Response{
		Data: data,
		Meta: meta{
			Version:   apiVersion,
			Timestamp: time.Now().String(),
		},
	})
}

func NewResponseWithDescription(c echo.Context, statusCode int, description string) error {
	return c.JSON(statusCode, Response{
		Meta: meta{
			Descipriton: description,
			Version:     apiVersion,
			Timestamp:   time.Now().String(),
		},
	})
}
