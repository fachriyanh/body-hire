package entity

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Title       string
	Description string
	AuthorID    uint
	PublisherID uint
}

type Author struct {
	gorm.Model
	Name string
}

type Publisher struct {
	gorm.Model
	Name string
}
