package routes

import (
	"github.com/devchrisar/go_gormAndrestAPI/controllers"
	"github.com/labstack/echo/v4"
)

func IndexRoute(router *echo.Echo) {
	router.GET("/", controllers.IndexHandler)
}
