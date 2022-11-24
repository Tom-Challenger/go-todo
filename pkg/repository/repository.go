package repository

import "github.com/jmoiron/sqlx"

type Autharization interface {
}

type todoList interface {
}

type todoItem interface {
}

type Repository struct {
	Autharization
	todoList
	todoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
