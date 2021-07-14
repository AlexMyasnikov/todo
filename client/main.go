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
			fmt.Printf("To do list:\n=================\nID:%d,\nName: %s,\nDescription: %s,\nStatus: %t\n",
				response.Id,
				response.Name,
				response.Description,
				response.Status)
		}
	case "update":
		{
			id, err := strconv.ParseInt(args[2], 10, 64)
			name := args[3]
			description := args[4]
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			request := &pb.UpdateRequest{Id: id, Name: name, Description: description}

			response, err := client.UpdateToDo(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			if response.Id != 0 {
				fmt.Printf("To do list:\n=================\nID:%d,\nName: %s,\nDescription: %s,\nStatus: %t\n",
					response.Id,
					response.Name,
					response.Description,
					response.Status)
			} else {
				fmt.Printf("To do with ID-%d doesn't exist\n", id)
			}
		}
	case "delete":
		{
			id, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			request := &pb.DeleteRequest{Id: id}

			response, err := client.DeleteToDo(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			if response.Exist == true {
				fmt.Printf("To do list with ID-%d has been deleted\n", id)
			} else {
				fmt.Printf("To do with ID-%d doesn't exist\n", id)
			}
		}

	case "get":
		{
			id, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			request := &pb.GetByIdRequest{Id: int64(id)}

			response, err := client.GetToDoById(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			if response.Id != 0 {
				fmt.Printf("To do list:\n=================\nID:%d,\nName: %s,\nDescription: %s,\nStatus: %t\n",
					response.Id,
					response.Name,
					response.Description,
					response.Status)
			} else {
				fmt.Printf("To do with ID-%d doesn't exist\n", id)
			}
		}
	case "getall":
		{
			request := &pb.GetAllRequest{}
			response, err := client.GetAllToDo(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			fmt.Printf("%s", "To do lists:\n")
			for i := range response.Todo {
				fmt.Printf("=================\nID:%d,\nName: %s,\nDescription: %s,\nStatus: %t\n",
					response.Todo[i].Id,
					response.Todo[i].Name,
					response.Todo[i].Description,
					response.Todo[i].Status)
			}
		}

	case "done":
		{
			id, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			request := &pb.MarkAsDoneRequest{Id: id}

			response, err := client.MarkAsDone(context.Background(), request)
			if err != nil {
				grpclog.Fatalf("%v", err)
			}
			if response.Id != 0 {
				fmt.Printf("To do list:\n=================\nID:%d,\nName: %s,\nDescription: %s,\nStatus: %t\n",
					response.Id,
					response.Name,
					response.Description,
					response.Status)
			} else {
				fmt.Printf("To do with ID-%d doesn't exist\n", id)
			}
		}
	default:
		fmt.Printf("This command doesn't exist\n")
	}

}
