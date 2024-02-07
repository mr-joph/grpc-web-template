package rpc

import (
	todoListRPC "core/pkg/handlers/todolist/rpc"

	"core/pkg/rpc/server"
)

func loadServices() {
	server.Register(&todoListRPC.TodoListServer{})
}
