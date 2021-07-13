package main

import (
	"context"
	"fmt"
	pb "github.com/ChuvashPeople/todo/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"os"
	"strconv"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	args := os.Args

	conn, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		grpclog.Fatalf("%v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewTodoClient(conn)

	request := args[1]
	switch request {
	case "create":
		{
			name := args[2]
			description := args[3]

			request := &pb.CreateRequest{
				Name:        name,
				Description: description,
				Status:      false,
			}
			response, err := client.CreateToDo(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			fmt.Println(response)

		}
	case "delete":
		{
			id, err := strconv.ParseInt(os.Args[2], 10, 64)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			request := &pb.DeleteRequest{Id: int64(id)}

			response, err := client.DeleteToDo(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			fmt.Printf("%s", response)
		}

	case "get":
		{
			id, err := strconv.ParseInt(os.Args[2], 10, 64)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			request := &pb.GetByIdRequest{Id: int64(id)}

			response, err := client.GetToDoById(context.Background(), request)
			if err != nil {
				fmt.Printf("%v", err)
			}
			fmt.Println(response)
		}

	}

}
