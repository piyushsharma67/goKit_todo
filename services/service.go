package service

type TodoServiceInterface interface {
	AddTodo(title, description string) (string, error)
	DeleteTodo(id int32) (string, error)
}

type todoService struct{}

func NewToDoService() TodoServiceInterface {
	return &todoService{}
}
