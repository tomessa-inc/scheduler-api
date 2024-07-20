package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetAbsenceRoutes(e *echo.Echo) {
	e.POST("/absence/page-index/:page-index/page-size/:page-size/:field/:order", controller.GetSchedule)
	e.POST("/absence/new", controller.SetAbsence)
	e.POST("/absence/available", controller.SetAvailable)

	//e.GET("/schedule/test", controller.Test)
}
