package main

import (
	"cms/controller"
	"cms/db/connection"
	middleware "cms/middleware/logging"
	"cms/repository"
	"cms/router"
	"cms/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error load .env file : %v\n", err.Error())
	}
}

func main() {
	log.Println("== run app ==")

	db := connection.ConnectDatabase()
	validate := validator.New()

	// inject repository
	userRepo := repository.NewUserRepository(db)
	accRepo := repository.NewAccountRepository(db)
	contentRepo := repository.NewContentRepository(db)
	likeRepo := repository.NewLikeRepository(db)
	dislikeRepo := repository.NewDislikeRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	// inject service
	accService := service.NewAccountService(accRepo, userRepo, validate)
	userService := service.NewUserService(userRepo, validate)
	contentService := service.NewContentService(contentRepo, accRepo, userRepo, validate)
	likeService := service.NewLikeService(likeRepo, accRepo, contentRepo, dislikeRepo, validate)
	dislikeService := service.NewDislikeService(likeRepo, dislikeRepo, accRepo, contentRepo, validate)
	reviewService := service.NewReviewService(reviewRepo, accRepo, contentRepo, validate)

	// inject controller
	accController := controller.NewAccountController(accService)
	userController := controller.NewUserController(userService)
	contentController := controller.NewContentController(contentService)
	likeController := controller.NewLikeController(likeService)
	dislikeController := controller.NewDislikeController(dislikeService)
	reviewController := controller.NewReviewController(reviewService)

	// inject router
	muxRouter := router.NewRouter(accController, userController, contentController, likeController, dislikeController, reviewController)

	// middleware
	loggerRepository := repository.NewLoggerRepository(db)
	loggerService := service.NewLoggerService(loggerRepository)
	loggerMiddleware := middleware.NewLoggerMiddleware(loggerService, muxRouter)

	// create server
	httpServer := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", os.Getenv("EXPOSE_PORT")),
		Handler: loggerMiddleware,
	}

	// run server
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("error when run server : %v\n", err.Error())
	}
}
