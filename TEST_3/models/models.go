package models

import (
	"github.com/google/uuid"
)

type Book struct {
	IDBook      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(255);not null"`
	IDAuthor    uuid.UUID `gorm:"type:uuid;not null"`
	IDPublisher uuid.UUID `gorm:"type:uuid;not null"`

	Author    *Author    `gorm:"foreignkey:IDAuthor;references:IDAuthor"`
	Publisher *Publisher `gorm:"foreignkey:IDPublisher;references:IDPublisher"`
}

type BookRequest struct {
	Name          string `json:"name"  binding:"required"`
	AuthorName    string `json:"author"  binding:"required"`
	PublisherName string `json:"publisher"  binding:"required"`
}

type BookResponse struct {
	Name          string `json:"name"  binding:"required"`
	AuthorName    string `json:"author"  binding:"required"`
	PublisherName string `json:"publisher"  binding:"required"`
}

type Author struct {
	IDAuthor uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string    `gorm:"type:varchar(255);not null"`
}

type Publisher struct {
	IDPublisher uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(255);not null"`
}

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string    `gorm:"type:varchar(255);not null"`
	Password string
}

type SignInInput struct {
	Name     string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
