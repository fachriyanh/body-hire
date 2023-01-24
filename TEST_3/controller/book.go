package controllers

import (
	"net/http"
	"strings"

	"test_3/TEST_3/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	DB *gorm.DB
}

func NewBookController(DB *gorm.DB) BookController {
	return BookController{DB}
}

func (bc *BookController) GetBookByName(ctx *gin.Context) {
	bookName := ctx.Param("name")

	var book []models.Book
	result := bc.DB.Find(&book, "name = ?", bookName)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "book not exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": book})
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var payload *models.BookRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var author models.Author
	resAuthor := bc.DB.Find(&author, "name = ?", payload.AuthorName)
	if resAuthor.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "author not exists"})
		return
	}

	var publisher models.Publisher
	resPublisher := bc.DB.Find(&publisher, "name = ?", payload.PublisherName)
	if resPublisher.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "publisher not exists"})
		return
	}

	newBook := models.Book{
		Name:        payload.Name,
		IDAuthor:    author.IDAuthor,
		IDPublisher: publisher.IDPublisher,
	}

	result := bc.DB.Create(&newBook)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": result})
}

func (pc *BookController) DeleteBook(ctx *gin.Context) {
	idBook := ctx.Param("id")

	result := pc.DB.Delete(&models.Book{}, "id = ?", idBook)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "book not exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (bc *BookController) UpdateBookByID(ctx *gin.Context) {
	idBook := ctx.Param("idBook")

	var payload *models.BookRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedBook models.BookRequest
	result := bc.DB.First(&updatedBook, "id = ?", idBook)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	var author models.Author
	resAuthor := bc.DB.Find(&author, "name = ?", payload.AuthorName)
	if resAuthor.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "author not exists"})
		return
	}

	var publisher models.Publisher
	resPublisher := bc.DB.Find(&publisher, "name = ?", payload.PublisherName)
	if resPublisher.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "publisher not exists"})
		return
	}

	updateBook := models.Book{
		Name:        payload.Name,
		IDAuthor:    author.IDAuthor,
		IDPublisher: publisher.IDPublisher,
	}

	bc.DB.Model(&updatedBook).Updates(updateBook)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedBook})
}
