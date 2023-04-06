package routes

import (
	"github.com/devchrisar/go_gormAndrestAPI/controllers"
	"github.com/labstack/echo/v4"
)

func TaskRoute(router *echo.Echo) {
	router.GET("/tasks", controllers.GetTasksHandler)
	router.POST("/tasks", controllers.CreateTasksHandler)
	router.GET("/tasks/:id", controllers.GetTaskHandler)
	router.DELETE("/tasks/:id", controllers.DeleteTasksHandler)
}
