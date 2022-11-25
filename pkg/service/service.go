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
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Autharization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autharization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
