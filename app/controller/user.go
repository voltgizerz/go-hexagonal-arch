package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/go-hexagonal-arch/app/entity"
	"github.com/voltgizerz/go-hexagonal-arch/app/port"
)

type UserController struct {
	userInteractor port.InMemoryUserInteractorInterface
}

func NewUserController(userInteractor port.InMemoryUserInteractorInterface) *UserController {
	return &UserController{
		userInteractor: userInteractor,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	user := &entity.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userInteractor.CreateUser(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUser(c *gin.Context) {
	user := &entity.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userInteractor.CreateUser(user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
