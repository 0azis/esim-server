package httprouter

import (
	"esim/internal/app"

	"github.com/labstack/echo/v4"
)

type Router struct {
	api *echo.Group
	app app.App
}

func NewRouter(app app.App) Router {
	api := app.Server.ApiRouter()
	return Router{api, app}
}

func (r Router) Init() {
	r.userRoutes()
}

func (r Router) userRoutes() {
	user := r.api.Group("/users")
	controller := NewUserControllers(r.app)

	user.POST("/auth/email", controller.AuthEmail)
	user.POST("/auth/telegram", controller.AuthTelegram)
}
