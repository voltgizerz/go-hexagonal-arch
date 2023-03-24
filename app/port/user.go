package port

import (
	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/go-hexagonal-arch/app/entity"
)

type InMemoryUserRepositoryInterface interface {
	Create(user *entity.User) error
	FindByID(id int) (*entity.User, error)
}

type InMemoryUserInteractorInterface interface {
	CreateUser(name string) (*entity.User, error)
	GetUser(c *gin.Context)
}
