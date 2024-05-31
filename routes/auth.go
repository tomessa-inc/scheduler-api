package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetAuthRoutes(e *echo.Echo) {
	e.POST("/auth/sign-in", controller.SignIn)
	e.POST("/auth/sign-out", controller.SignOut)
}
