package router

import (
	"github.com/gin-gonic/gin"
	"training/app/handler"
)

type UserRouter struct {
	UserHandler *handler.UserHandler
}

// create function provider
func NewUserRouter(userHandler *handler.UserHandler) *UserRouter {
	return &UserRouter{
		UserHandler: userHandler,
	}
}

func (u *UserRouter) CreateRoutes(r *gin.RouterGroup) *gin.RouterGroup {
	route := r.Group("")
	route.POST("/user", u.UserHandler.SignUp)
	route.POST("/login", u.UserHandler.Login)

	return route
}
