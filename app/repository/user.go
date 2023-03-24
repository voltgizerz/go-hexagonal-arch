package repository

import (
	"github.com/voltgizerz/go-hexagonal-arch/app/entity"
	"github.com/voltgizerz/go-hexagonal-arch/app/port"
)

type InMemoryUserRepository struct {
	users map[int]*entity.User
}

func NewInMemoryUserRepository() port.InMemoryUserRepositoryInterface {
	return &InMemoryUserRepository{}
}

func (ir *InMemoryUserRepository) FindByID(id int) (*entity.User, error) {
	if user, ok := ir.users[id]; ok {
		return user, nil
	}
	return nil, nil
}

func (ir *InMemoryUserRepository) Create(user *entity.User) error {
	ir.users[user.ID] = user
	return nil
}
