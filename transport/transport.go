package transport

import (
	"context"
	"gokit_todo/endpoint"
	"gokit_todo/todo"

	"github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	addTodo    grpc.Handler
	deleteTodo grpc.Handler
	todo.UnimplementedTodoServiceServer
}

func NewGRPCServer(endpoints endpoint.Endpoints) todo.TodoServiceServer {
	return &grpcServer{
		addTodo: grpc.NewServer(
			endpoints.AddTodoEndpoint,
			decodeGRPCAddTodoRequest,
			encodeGRPCAddTodoResponse,
		),
		deleteTodo: grpc.NewServer(
			endpoints.DeleteTodoEndpoint,
			decodeGRPCDeleteTodoRequest,
			encodeGRPCDeleteTodoResponse,
		),
	}
}

func (s *grpcServer) AddTodo(ctx context.Context, req *todo.AddTodoRequest) (*todo.AddTodoResponse, error) {
	_, resp, err := s.addTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*todo.AddTodoResponse), nil
}

func (s *grpcServer) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest) (*todo.DeleteTodoResponse, error) {
	_, resp, err := s.deleteTodo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*todo.DeleteTodoResponse), nil
}

// Decoders
func decodeGRPCAddTodoRequest(_ context.Context, req interface{}) (interface{}, error) {
	r := req.(*todo.AddTodoRequest)
	return endpoint.AddTodoRequest{Title: r.Title, Description: r.Description}, nil
}

func decodeGRPCDeleteTodoRequest(_ context.Context, req interface{}) (interface{}, error) {
	r := req.(*todo.DeleteTodoRequest)
	return endpoint.DeleteTodoRequest{ID: r.Id}, nil
}

// Encoders
func encodeGRPCAddTodoResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r := resp.(endpoint.AddTodoResponse)
	return &todo.AddTodoResponse{Status: r.Status, Error: r.Err.Error()}, nil
}

func encodeGRPCDeleteTodoResponse(_ context.Context, resp interface{}) (interface{}, error) {
	r := resp.(endpoint.DeleteTodoResponse)
	return &todo.DeleteTodoResponse{Status: r.Status, Error: r.Err.Error()}, nil
}
