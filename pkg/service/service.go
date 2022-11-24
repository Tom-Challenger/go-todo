package service

import (
	"github.com/Tom-Challenger/go-todo"
	"github.com/Tom-Challenger/go-todo/pkg/repository"
)

/* Интерфейсы определяются в соответствии с доменной зоной.
Важно выделить в бизнес логике основные части */

type Autharization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type todoList interface {
}

type todoItem interface {
}

type Service struct {
	Autharization
	todoList
	todoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autharization: NewAuthService(repos.Authorization),
	}
}
