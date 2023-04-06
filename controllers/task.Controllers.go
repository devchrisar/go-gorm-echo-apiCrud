package controllers

import (
	"encoding/json"
	"github.com/devchrisar/go_gormAndrestAPI/db"
	"github.com/devchrisar/go_gormAndrestAPI/models"
	"github.com/labstack/echo/v4"
)

func GetTasksHandler(c echo.Context) error {
	var tasks []models.Task
	db.DB.Find(&tasks)
	return json.NewEncoder(c.Response().Writer).Encode(&tasks)
}

func CreateTasksHandler(c echo.Context) error {
	var task models.Task
	json.NewDecoder(c.Request().Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		c.String(400, err.Error())
		return nil
	}
	return json.NewEncoder(c.Response().Writer).Encode(&task)
}

func GetTaskHandler(c echo.Context) error {
	var task models.Task
	params := c.Param("id")
	db.DB.First(&task, params)
	if task.ID == 0 {
		c.String(404, "Task not found")
		return nil
	}
	return json.NewEncoder(c.Response().Writer).Encode(&task)
}

func DeleteTasksHandler(c echo.Context) error {
	var task models.Task
	params := c.Param("id")
	db.DB.First(&task, params)
	if task.ID == 0 {
		c.String(404, "Task not found")
		return nil
	}
	db.DB.Unscoped().Delete(&task)
	return c.String(204, "")
}
