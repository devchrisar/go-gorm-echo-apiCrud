package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DSN = "host=localhost user=**USER** password=**PASSWORD** dbname=**DATABASE** port=**PORT**"
var DB *gorm.DB

func DBconn() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection established")
	}
}
