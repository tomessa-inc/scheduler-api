package routes

import (
	"net/http"
	"regexp"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	SetUserRoutes(e)
	SetUsherGroupRoutes(e)
	SetWeekRoutes(e)
	SetAuthRoutes(e)
	SetScheduleRoutes(e)
}

func paramValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		paramKey := c.ParamNames()
		paramValue := c.ParamValues()

		r := regexp.MustCompile("^[a-z]+$")

		for k, v := range paramValue {
			if !r.MatchString(v) {
				return c.JSON(http.StatusBadRequest, "param ("+paramKey[k]+") is not a number")
			}
		}

		return next(c)
	}

}

//func HandleRequestsGallery() {

//	myRouter := mux.NewRouter().StrictSlash(true)/
//	//test := controller.MainGallery
//	myRouter.HandleFunc("/gallery", controller.MainGallery(http.ResponseWriter, http.Request))
///	log.Fatal(http.ListenAndServe(":10000", nil))
//}
