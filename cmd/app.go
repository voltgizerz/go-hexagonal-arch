package main

import (
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/go-hexagonal-arch/app/api"
	"github.com/voltgizerz/go-hexagonal-arch/app/config"
	"github.com/voltgizerz/go-hexagonal-arch/app/controller"
	"github.com/voltgizerz/go-hexagonal-arch/app/interactor"
	"github.com/voltgizerz/go-hexagonal-arch/app/repository"
)

var (
	log = config.SetupLog()
)

func main() {
	config.LoadENV()

	// Initialize repository layer
	userRepository := repository.NewInMemoryUserRepository()

	// Initialize interactor layer
	userInteractor := interactor.NewUserInteractor(userRepository)

	// Initialize controller layer
	userController := controller.NewUserController(userInteractor)

	// Init Router API
	router := &api.Router{
		Router:         gin.Default(),
		UserController: userController,
	}

	router.StartServer()
}
