package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	db2 "training/app/db"
	"training/app/handler"
	"training/app/repository"
	"training/app/router"
	"training/app/service"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error load env:", err.Error())
	}
}

func main() {
	// register db layer
	db := db2.NewConnection()

	// register repository layer
	bookRepo := repository.NewBookRepository(db)
	userRepo := repository.NewUserRepository(db)

	// register service layer
	bookService := service.NewBookService(bookRepo, userRepo)
	userService := service.NewUserRepository(userRepo)

	// register handler layer
	bookHandler := handler.NewBookHander(bookService)
	userHandler := handler.NewUserHandler(userService)

	// register router layer
	bookRouter := router.NewBookRouter(bookHandler)
	userRouter := router.NewUserRouter(userHandler)

	// create router
	router := gin.Default()
	routerv1 := router.Group("/api/v1")

	// == book ==
	bookRouter.CreateRoutes(routerv1)

	// == user ==
	userRouter.CreateRoutes(routerv1)

	// run
	router.Run(":5001")
}
