package model

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Pages       int     `json:"pages" gorm:"check>100"`
	Author      *Author `json:"author" gorm:"embedded;embeddedPrefix:author_"`
}

type Message struct {
	Message string `json:"message"`
}
