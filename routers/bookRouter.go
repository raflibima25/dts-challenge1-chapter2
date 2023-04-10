package routers

import (
	"challenge-2-chapter-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books/", controllers.GetBookAll)

	router.GET("/books/:bookID", controllers.GetBookId)

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:bookID", controllers.UpdateBook)

	router.DELETE("books/:bookID", controllers.DeleteBook)

	return router
}
