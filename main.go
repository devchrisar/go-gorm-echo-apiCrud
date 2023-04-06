package main

import (
	"github.com/devchrisar/go_gormAndrestAPI/db"
	"github.com/devchrisar/go_gormAndrestAPI/models"
	"github.com/devchrisar/go_gormAndrestAPI/routes"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	db.DBconn()
	err := db.DB.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.DB.AutoMigrate(models.Task{})
	if err != nil {
		log.Fatal(err)
	}

	router := echo.New()
	routes.IndexRoute(router)
	routes.UserRoute(router)
	routes.TaskRoute(router)

	router.Logger.Fatal(router.Start(":" + port))
}
