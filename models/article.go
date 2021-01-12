// models/article.go

package models

type Article struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Author      string `json:"author"`
}