package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Errors struct {
	Values []err `json:"errors"`
}

type err struct {
	StatusCode  int    `json:"code"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewErrors(c echo.Context, statusCode int, errors ...err) error {
	var errs Errors

	for _, e := range errors {
		errs.Values = append(errs.Values, e)
	}

	return c.JSON(statusCode, errs)
}

func Error(statusCode int, title, description string) err {
	return err{
		StatusCode:  statusCode,
		Status:      http.StatusText(statusCode),
		Title:       title,
		Description: description,
	}
}
