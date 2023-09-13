package router

import (
	"cms/controller"
	"cms/helper"
	"cms/middleware/auth"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// create function mux
func NewRouter(accController *controller.AccountController, userController *controller.UserController, contentController *controller.ContentController, likeController *controller.LikeController, dislikeController *controller.DislikeController, reviewController *controller.ReviewController) *mux.Router {
	r := mux.NewRouter()

	namsReo := "reo"
	namsReo = "rona"
	nilai := 100

	var nilai1 = int(20)

	fmt.Printf(namsReo)
	fmt.Print(nilai, nilai1)

	// create middleware
	authMiddleware := auth.NewAuthMiddleware(r)

	// ovveride handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.ResponseError(w, http.StatusNotFound, "endpoint not exist")
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.ResponseError(w, http.StatusMethodNotAllowed, "method not allowed")
	})

	// buat endpoint
	// test
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		helper.ResponseSuccess(w, "success test cms service", map[string]any{
			"welcome":   fmt.Sprintf("welcome to cms app node-%v", os.Getenv("APP_NODE")),
			"timestamp": time.Time(time.Now()).Format("2006-01-02 15:04:05"),
		})
	}).Methods(http.MethodGet)

	// account
	r.HandleFunc("/api/account", accController.CreateAccount).Methods(http.MethodPost)
	r.Handle("/api/account-check", authMiddleware.Auth(http.HandlerFunc(accController.CheckAccount))).Methods(http.MethodGet)
	r.HandleFunc("/api/otp", accController.RequestOTP).Methods(http.MethodPost)

	// login
	r.HandleFunc("/api/login", accController.Login).Methods(http.MethodPost)

	// forgot password
	r.HandleFunc("/api/forgot-password", accController.ForgotPassword).Methods(http.MethodPost)

	// change password
	r.Handle("/api/change-password", authMiddleware.Auth(http.HandlerFunc(accController.ChangePassword))).Methods(http.MethodPut)

	// get user by email
	r.Handle("/api/user", authMiddleware.Auth(http.HandlerFunc(userController.GetByEmail))).Methods(http.MethodGet)

	// get all users
	r.HandleFunc("/api/users", userController.GetAll).Methods(http.MethodGet)

	// create content
	r.Handle("/api/content", authMiddleware.Auth(http.HandlerFunc(contentController.Insert))).Methods(http.MethodPost)

	// get content by ID
	r.Handle("/api/content/{id}", authMiddleware.Auth(http.HandlerFunc(contentController.GetById))).Methods(http.MethodGet)

	// Edit / Update content by ID
	r.Handle("/api/content", authMiddleware.Auth(http.HandlerFunc(contentController.Update))).Methods(http.MethodPut)

	// like content
	r.Handle("/api/content/like", authMiddleware.Auth(http.HandlerFunc(likeController.Like))).Methods(http.MethodPost)

	// unlike content
	r.Handle("/api/content/unlike", authMiddleware.Auth(http.HandlerFunc(likeController.Unlike))).Methods(http.MethodPost)

	// dislike content
	r.Handle("/api/content/dislike", authMiddleware.Auth(http.HandlerFunc(dislikeController.Dislike))).Methods(http.MethodPost)

	// undislike content
	r.Handle("/api/content/undislike", authMiddleware.Auth(http.HandlerFunc(dislikeController.Undislike))).Methods(http.MethodPost)

	// get all contents by creator (username)
	r.Handle("/api/{username}/content", authMiddleware.Auth(http.HandlerFunc(contentController.GetByUsername))).Methods(http.MethodGet)

	// get all contents
	r.Handle("/api/contents", authMiddleware.Auth(http.HandlerFunc(contentController.GetAll))).Methods(http.MethodGet)

	// get likes by email/username -> {username}/like
	r.Handle("/api/{username}/like", authMiddleware.Auth(http.HandlerFunc(likeController.GetContentLikedByUsername))).Methods(http.MethodGet)

	// get dislike by username -> {username}/dislike
	r.Handle("/api/{username}/dislike", authMiddleware.Auth(http.HandlerFunc(dislikeController.GetContentDislikedByUsername))).Methods(http.MethodGet)

	// create reviews
	r.Handle("/api/review", authMiddleware.Auth(http.HandlerFunc(reviewController.CreateReview))).Methods(http.MethodPost)

	// delete review
	r.Handle("/api/review", authMiddleware.Auth(http.HandlerFunc(reviewController.Delete))).Methods(http.MethodDelete)

	// return
	return r
}
