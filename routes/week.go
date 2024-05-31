package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetWeekRoutes(e *echo.Echo) {
	e.POST("/week", controller.AddWeek)
	e.GET("/week/details/:weekId", controller.GetDetails)

	// e.GET("/gallery/category/:category", controller.GalleryByCategory, paramValidation)
	// e.GET("/gallery/tag/:tag", controller.GalleryByTag, paramValidation)
}
