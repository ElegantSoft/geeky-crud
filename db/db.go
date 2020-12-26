package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitConnection() *gorm.DB {
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second, // Slow SQL threshold
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      false,       // Disable color
	//	},
	//)
	log.Print("will establish")
	conn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	log.Printf("will establish %+v", conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn
}
