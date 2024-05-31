package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetUserRoutes(e *echo.Echo) {
	e.POST("/user", controller.AddUser)
	e.GET("/user/page-index/:page-index/page-size/:page-size/:field/:order", controller.GetUsers)
	e.GET("/user/page-index/:page-index/page-size/:page-size/:field/:order/:prefix", controller.GetUsersByPrefix)

	e.GET("/user/id/:id", controller.GetUserById)
	e.PUT("/user/id/:id", controller.UpdateUser)

	e.GET("/user/get-user/:token", controller.GetUserByToken)
	e.GET("/user/auth/:token", controller.GetUserAuthByToken)

	e.POST("/user/reset-password", controller.ResetPassword)
}
