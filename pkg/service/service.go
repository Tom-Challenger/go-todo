package service

import "github.com/Tom-Challenger/go-todo/pkg/repository"

/* Интерфейсы определяются в соответствии с доменной зоной.
Важно выделить в бизнес логике основные части */

type Autharization interface {

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
	return &Service{}
}