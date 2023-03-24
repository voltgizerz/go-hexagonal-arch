package interactor

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/voltgizerz/go-hexagonal-arch/app/entity"
	"github.com/voltgizerz/go-hexagonal-arch/app/port"
	"github.com/voltgizerz/go-hexagonal-arch/app/utils"
)

type UserInteractor struct {
	userRepository port.InMemoryUserRepositoryInterface
}

func NewUserInteractor(userRepository port.InMemoryUserRepositoryInterface) port.InMemoryUserInteractorInterface {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

func (ui *UserInteractor) CreateUser(name string) (*entity.User, error) {
	user := &entity.User{Name: name}
	err := ui.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ui *UserInteractor) GetUser(c *gin.Context) {
	id := utils.StringToInt(c.Param("id"))
	user, err := ui.userRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
