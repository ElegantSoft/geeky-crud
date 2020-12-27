package crud

import (
	"github.com/ElegantSoft/geeky-crud/db"
	"github.com/ElegantSoft/geeky-crud/models"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

func TestCrud_Find(t *testing.T) {
	db.InitConnection()

	db.Conn.AutoMigrate(&models.Post{})
	service := NewCrudService(db.Conn)
	post := []models.Post{
		{
			Title:    "test1",
			Content:  "asddsaasddsa",
			Comments: nil,
		},
		{
			Title:    "test1",
			Content:  "test1",
			Comments: nil,
		},
	}
	db.Conn.Create(&post)
	var a models.Post
	service.Find(
		FindQuery{
			Q: Query{
				"content": {
					"$like": "asd",
				},
				"title": {
					"$eq":  "test1",
					"$not": "asd",
				},
			},
		},
		&a)
	log.Printf("a -> %+v", a)

	assert.Equal(t, a.Title, "test1")
}
