package main

import (
	"github.com/ElegantSoft/geeky-crud/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	server := gin.Default()
	//conn, err := gorm.Open(sqlite.Open("as.db"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	// Migrate the schema
	if err := conn.AutoMigrate(&models.Comment{}); err != nil {
		log.Fatal(err.Error())
	}
	// Migrate the schema
	//if err := db.Conn.AutoMigrate(&models.Post{}); err != nil {
	//	log.Fatal(err.Error())
	//}

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
