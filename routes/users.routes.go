package routes

import (
	"github.com/devchrisar/go_gormAndrestAPI/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoute(router *echo.Echo) {
	router.GET("/users", controllers.GetUsersHandler)
	router.GET("/users/:id", controllers.GetUserHandler)
	router.POST("/users", controllers.PostUserHandler)
	router.DELETE("/users/:id", controllers.DeleteUserHandler)
}
