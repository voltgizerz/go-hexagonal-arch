package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/go-hexagonal-arch/app/config"
	"github.com/voltgizerz/go-hexagonal-arch/app/controller"
)

var (
	log         = config.SetupLog()
	defaultPort = ":8080"
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

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = defaultPort
	}

	// Start server
	if err := r.Router.Run(port); err != nil {
		log.Fatal(err)
	}
	log.Info("Server started on localhost" + port)
}
