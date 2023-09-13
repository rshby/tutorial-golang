package testing

import (
	"cms/controller"
	middleware "cms/middleware/logging"
	"cms/repository"
	"cms/router"
	"cms/service"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

// create variabel refer to router
var routerTest = NewRouterTesting()

// create function to connect with database
func ConnectDBTesting() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/cms_test?parseTime=true")

	// jika ada kesalahan ketika connect
	if err != nil {
		log.Fatalf("error connection database : %v\n", err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	db.SetConnMaxIdleTime(30 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)

	return db
}

// create router for testing
func NewRouterTesting() http.Handler {
	db := ConnectDBTesting()
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

	return loggerMiddleware
}
