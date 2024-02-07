package rpc

import (
	"context"
	"core/pkg/helper/rpcutils"
	"core/pkg/proto/pb"
	"core/pkg/storage"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoListServer struct {
	pb.UnimplementedTodoListServiceServer
}

func (s *TodoListServer) GetTodos(
	ctx context.Context,
	req *emptypb.Empty,
) (*pb.GetTodosOut, error) {
	todos := storage.GetAll()
	return &pb.GetTodosOut{
		Todos: todos,
	}, nil
}

func (s *TodoListServer) NewTodo(
	ctx context.Context,
	req *pb.NewTodoIn,
) (*pb.NewTodoOut, error) {
	todo := &pb.Todo{}
	todo.Id = uuid.New().String()
	todo.Content = req.Content
	todo.Completed = false

	storage.Add(todo)

	return &pb.NewTodoOut{
		Todo: todo,
	}, nil
}

func (s *TodoListServer) EditTodo(
	ctx context.Context,
	todo *pb.Todo,
) (*pb.EditTodoOut, error) {
	editedTodo, err := storage.Update(todo)
	if err != nil {
		rpcutils.SendNotFoundError("Todo doesn't exist")
	}

	return &pb.EditTodoOut{
		Todo: editedTodo,
	}, nil
}

func (s *TodoListServer) CheckTodo(
	ctx context.Context,
	req *pb.CheckTodoIn,
) (*pb.CheckTodoOut, error) {
	todo, _, err := storage.GetById(req.Id)
	if err != nil {
		rpcutils.SendNotFoundError("Todo doesn't exist")
	}

	todo.Completed = req.Completed
	_, err = storage.Update(todo)
	if err != nil {
		rpcutils.SendAbortError(err)
	}

	return &pb.CheckTodoOut{
		Todo: todo,
	}, nil
}

func (s *TodoListServer) DeleteTodo(
	ctx context.Context,
	req *pb.DeleteTodoIn,
) (*emptypb.Empty, error) {
	todo, _, err := storage.GetById(req.Id)
	if err != nil {
		rpcutils.SendNotFoundError("Todo doesn't exist")
	}

	storage.Remove(todo)

	return &emptypb.Empty{}, nil
}

func (s *TodoListServer) CompleteAll(
	ctx context.Context,
	req *emptypb.Empty,
) (*emptypb.Empty, error) {
	todos := storage.GetAll()
	for _, todo := range todos {
		todo.Completed = true
		storage.Update(todo)
	}

	return &emptypb.Empty{}, nil
}

func (s *TodoListServer) ClearCompleted(
	ctx context.Context,
	req *emptypb.Empty,
) (*emptypb.Empty, error) {
	todos := storage.GetAll()
	toRemove := []*pb.Todo{}

	for _, todo := range todos {
		if todo.Completed {
			toRemove = append(toRemove, todo)
		}
	}

	for _, todo := range toRemove {
		storage.Remove(todo)
	}

	return &emptypb.Empty{}, nil
}

func (s *TodoListServer) RegisterOnServer(server *grpc.Server) {
	pb.RegisterTodoListServiceServer(server, s)
}
