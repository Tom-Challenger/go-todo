package repository

import (
	"github.com/Tom-Challenger/go-todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type todoList interface {
}

type todoItem interface {
}

type Repository struct {
	Authorization
	todoList
	todoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
