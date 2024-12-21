package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID          int    `json:"ID"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var books = []Book{}

func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, &books)
}

func GetBooksByID(c *gin.Context) {
	id := c.Param("ID")
	for _, book := range books {
		if strconv.Itoa(book.ID) == id {
			c.JSON(http.StatusOK, &book)
		}
	}
}

func AddBook(c *gin.Context) {
	book := Book{}

	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, Book{ID: book.ID, Name: book.Name, Author: book.Author, Description: book.Description})
	c.JSON(http.StatusOK, &books)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("ID")
	for i, book := range books {
		if strconv.Itoa(book.ID) == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
}

func UpdateBook(c *gin.Context) {
	id := c.Param("ID")
	bookNewData := Book{}
	err := c.ShouldBindBodyWithJSON(&bookNewData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	for i, book := range books {
		if strconv.Itoa(book.ID) == id {
			books[i].Name = bookNewData.Name
			books[i].Author = bookNewData.Author
			books[i].Description = book.Description
		}
	}
	c.JSON(http.StatusOK, &books)
}
