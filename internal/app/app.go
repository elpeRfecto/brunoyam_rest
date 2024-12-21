package app

import (
	"brunoyam_rest/internal/transport"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/books", transport.GetBooks)
	r.POST("/books/:id", transport.AddBook)
	r.GET("/books/:id", transport.GetBooksByID)
	r.DELETE("/books/:id", transport.DeleteBook)
	r.PUT("/books/:id", transport.UpdateBook)
	if err := r.Run(); err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")r.Run()
}
