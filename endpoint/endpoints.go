package endpoint

import (
	"context"
	service "gokit_todo/services"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AddTodoEndpoint    endpoint.Endpoint
	DeleteTodoEndpoint endpoint.Endpoint
}

type AddTodoRequest struct {
	Title       string
	Description string
}

type AddTodoResponse struct {
	Status string
	Err    error
}

func MakeAddTodoEndpoint(svc service.TodoServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddTodoRequest)
		status, err := svc.AddTodo(req.Title, req.Description)
		return AddTodoResponse{Status: status, Err: err}, nil
	}
}

type DeleteTodoRequest struct {
	ID int32
}

type DeleteTodoResponse struct {
	Status string
	Err    error
}

// MakeDeleteTodoEndpoint converts DeleteTodo into an endpoint
func MakeDeleteTodoEndpoint(svc service.TodoServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteTodoRequest)
		status, err := svc.DeleteTodo(req.ID)
		return DeleteTodoResponse{Status: status, Err: err}, nil
	}
}
