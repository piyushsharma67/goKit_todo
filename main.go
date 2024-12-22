package main

import (
	"gokit_todo/endpoint"
	service "gokit_todo/services"
	"gokit_todo/todo"
	"gokit_todo/transport"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	service:=service.NewToDoService()

	addTodoEndpoint:=endpoint.MakeAddTodoEndpoint(service)
	deleteTodoEndpoint:=endpoint.MakeDeleteTodoEndpoint(service)

	endpoints := endpoint.Endpoints{
		AddTodoEndpoint:    addTodoEndpoint,
		DeleteTodoEndpoint: deleteTodoEndpoint,
	}

	grpcServer := transport.NewGRPCServer(endpoints)

	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Failed to listen on port 5000: %v", err)
	}

	server := grpc.NewServer()
	todo.RegisterTodoServiceServer(server, grpcServer)

	log.Println("gRPC server is running on port 5000...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}

}
