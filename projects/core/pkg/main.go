package main

import (
	"fmt"

	"core/pkg/helper/config"
	"core/pkg/rpc"
)

func main() {
	// Ready, Steady, Go!! 🔥
	fmt.Println("Hello gRPC web template 👋")
	fmt.Println(config.Env.Info.Name)
	fmt.Println(config.Env.Info.Version)

	// starts the gRPC server
	rpc.StartServer()
}
