package main

import (
	"context"
	"github.com/ChuvashPeople/todo/data"
	pb "github.com/ChuvashPeople/todo/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		grpclog.Fatalf("%v", err)
	}

	db := data.FakeDb{}
	server := NewToDoServer(&db)
	grpcServer := grpc.NewServer()

	pb.RegisterTodoServer(grpcServer, server)
	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("%v", err)
	}

}

type Server struct {
	db *data.FakeDb
}

func NewToDoServer(db *data.FakeDb) pb.TodoServer {
	return &Server{db: db}
}

func (s *Server) CreateToDo(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	todo := s.db.Create(request)
	return &pb.CreateResponse{Id: todo.Id, Name: todo.Name, Description: todo.Description, Status: todo.Status}, nil
}

func (s *Server) DeleteToDo(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	panic("implement me")
}

func (s *Server) GetToDoById(ctx context.Context, request *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	todo := s.db.Get(request)
	return &pb.GetByIdResponse{Id: todo.Id, Name: todo.Name, Description: todo.Description, Status: todo.Status}, nil
}

func (s *Server) GetAllToDo(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	panic("implement me")
}
