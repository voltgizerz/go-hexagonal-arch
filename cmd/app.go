package main

import (
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/go-hexagonal-arch/app/config/api"
	"github.com/voltgizerz/go-hexagonal-arch/app/controller"
	"github.com/voltgizerz/go-hexagonal-arch/app/interactor"
	"github.com/voltgizerz/go-hexagonal-arch/app/repository"
)

func main() {
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
