package repository


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

func NewRepository() *Repository {
	return &Repository{}
}