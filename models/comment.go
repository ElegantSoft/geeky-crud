package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text string `json:"text"`
	PostID uint `json:"post_id"`
}
