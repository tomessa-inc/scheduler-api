package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetScheduleRoutes(e *echo.Echo) {
	e.POST("/schedule/page-index/:page-index/page-size/:page-size/:field/:order", controller.GetSchedule)
	e.POST("/schedule/unavailable", controller.SetUnAvailable)
	e.POST("/schedule/available", controller.SetAvailable)

	//e.GET("/schedule/test", controller.Test)
}
