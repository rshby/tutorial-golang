package router

import (
	"github.com/gin-gonic/gin"
	"training/app/handler"
	"training/app/middleware"
)

type BookRouter struct {
	BookHandler *handler.BookHandler
}

// function provider
func NewBookRouter(bookHandler *handler.BookHandler) *BookRouter {
	return &BookRouter{
		BookHandler: bookHandler,
	}
}

// method to create routes
func (b *BookRouter) CreateRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	authMiddleware := middleware.NewAuthMiddleware()
	router := r.Group("")
	router.Use(authMiddleware.Auth())

	router.GET("/books", b.BookHandler.GetAllBooks)
	router.GET("/book/:id", b.BookHandler.GetBookById2)
	router.POST("/book", b.BookHandler.InsertBook)
	router.DELETE("/book/:id", b.BookHandler.DeleteBookById)
	router.PUT("/book", b.BookHandler.UpdateById)
	router.GET("/user/:id/books", b.BookHandler.GetByUserId)

	return r
}
