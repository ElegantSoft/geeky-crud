package main

import (
	"github.com/ElegantSoft/geeky-crud/db"
	"github.com/ElegantSoft/geeky-crud/models"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	db.InitConnection()
	server := gin.Default()
	// Migrate the schema
	if err := db.Conn.AutoMigrate(&models.Comment{}); err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Conn.AutoMigrate(&models.Post{}); err != nil {
		log.Fatal(err.Error())
	}

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
