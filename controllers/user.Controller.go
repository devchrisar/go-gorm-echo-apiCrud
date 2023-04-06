package controllers

import (
	"encoding/json"
	"github.com/devchrisar/go_gormAndrestAPI/db"
	"github.com/devchrisar/go_gormAndrestAPI/models"
	"github.com/labstack/echo/v4"
)

func GetUsersHandler(c echo.Context) error {
	var user []models.User
	db.DB.Find(&user)
	return json.NewEncoder(c.Response().Writer).Encode(&user)
}

func GetUserHandler(c echo.Context) error {
	var user models.User
	params := c.Param("id")
	db.DB.First(&user, params)
	if user.ID == 0 {
		c.String(404, "User not found")
		return nil
	}
	db.DB.Model(&user).Association("Task").Find(&user.Task)
	return json.NewEncoder(c.Response().Writer).Encode(&user)
}

func PostUserHandler(c echo.Context) error {
	var user models.User
	json.NewDecoder(c.Request().Body).Decode(&user)
	CreatedUser := db.DB.Create(&user)
	err := CreatedUser.Error
	if err != nil {
		return c.String(400, err.Error())
	}
	return json.NewEncoder(c.Response().Writer).Encode(user)
}

func DeleteUserHandler(c echo.Context) error {
	var user models.User
	params := c.Param("id")
	db.DB.First(&user, params)
	if user.ID == 0 {
		c.String(404, "User not found")
		return nil
	}
	db.DB.Unscoped().Delete(&user)
	return c.String(200, "")
}
