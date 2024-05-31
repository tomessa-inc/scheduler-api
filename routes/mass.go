package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetMassRoutes(e *echo.Echo) {
	e.POST("/mass", controller.AddUser)
	e.GET("/user/page-index/:page-index/page-size/:page-size/:field/:order", controller.GetUsers)
	// e.GET("/gallery/category/:category", controller.GalleryByCategory, paramValidation)
	// e.GET("/gallery/tag/:tag", controller.GalleryByTag, paramValidation)
}
