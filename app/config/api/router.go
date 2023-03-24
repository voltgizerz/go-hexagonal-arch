package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/go-hexagonal-arch/app/controller"
)

type Router struct {
	Router         *gin.Engine
	UserController *controller.UserController
}

func (r *Router) InitUserRouter() {
	r.Router.POST("/users", r.UserController.CreateUser)
	r.Router.GET("/users/:id", r.UserController.GetUser)
}

func (r *Router) StartServer() {
	// Initialize route
	r.InitUserRouter()

	// Start server
	if err := r.Router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
