package crud

import (
	"github.com/ElegantSoft/geeky-crud/db"
	"github.com/ElegantSoft/geeky-crud/models"
	"log"
	"testing"
)

func TestCrud_Find(t *testing.T) {
	db.InitConnection()

	db.Conn.AutoMigrate(&models.Post{})
	service := NewCrudService(db.Conn)
	post := models.Post{
		Title:    "test1",
		Content:  "test1",
		Comments: nil,
	}
	db.Conn.Create(&post)
	var a models.Post
	service.Find(
		findQuery{map[string]map[string]string{
			"title": {
				"$eq": "test1",
			},
			"content": {
				"$like": "tes",
			},
		}, []string{}},
		&a)
	log.Printf("a -> %+v", a)

	if a.ID != 1 {
		t.Fatal("failed")
	}
}
