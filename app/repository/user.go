package repository

import (
	"errors"

	"github.com/voltgizerz/go-hexagonal-arch/app/entity"
	"github.com/voltgizerz/go-hexagonal-arch/app/port"
)

type InMemoryUserRepository struct {
	users map[int]*entity.User
}

func NewInMemoryUserRepository() port.InMemoryUserRepositoryInterface {
	// Create some sample user data
	users := map[int]*entity.User{
		1: &entity.User{
			ID:   1,
			Name: "Alice",
		},
		2: &entity.User{
			ID:   2,
			Name: "Bob",
		},
		3: &entity.User{
			ID:   3,
			Name: "Charlie",
		},
	}

	return &InMemoryUserRepository{
		users: users,
	}
}

func (ir *InMemoryUserRepository) FindByID(id int) (*entity.User, error) {
	if user, ok := ir.users[id]; ok {
		return user, nil
	}
	return nil, errors.New("UserNotFound")
}

func (ir *InMemoryUserRepository) Create(user *entity.User) error {
	ir.users[user.ID] = user
	return nil
}
