package crud

import (
	"github.com/ElegantSoft/geeky-crud/db"
	"github.com/ElegantSoft/geeky-crud/models"
	"github.com/go-playground/assert/v2"
	"gorm.io/gorm"
	"log"
	"testing"
)

type Comment struct {
	gorm.Model
	Text string `json:"text"`
	PostID uint `json:"post_id"`
}

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	Comments []Comment `json:"comments"`
}

func TestCrud_Find(t *testing.T) {
	db.InitConnection()

	db.Conn.AutoMigrate(&Post{})
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
			Joins: []string{"comments"},
		},
		&a)
	log.Printf("a -> %+v", a)

	assert.Equal(t, a.Title, "test1")
}

func TestCrud_FindWithPreload(t *testing.T) {
	db.InitConnection()

	db.Conn.AutoMigrate(&Post{})
	db.Conn.AutoMigrate(&Comment{})
	db.Conn.Exec("DELETE from comments")
	db.Conn.Exec("DELETE from posts")
	service := NewCrudService(db.Conn)
	post := []models.Post{
		{
			Title:    "test1",
			Content:  "asddsaasddsa",
		},
		{
			Title:    "test1",
			Content:  "test1",
		},
	}
	db.Conn.Create(&post)

	comment := &Comment{
		Text:   "test",
		PostID: post[0].ID,
	}

	db.Conn.Create(&comment)

	log.Printf("comment -> %+v", comment)

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
			Joins: []string{"comments"},
		},
		&a)
	log.Printf("a -> %+v", a)

	assert.Equal(t, a.Title, post[0].Title)
	assert.Equal(t, a.Comments[0].ID, comment.ID)
}
