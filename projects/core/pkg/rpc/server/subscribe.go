package server

import "google.golang.org/grpc"

// finishing "er" for interfaces is a common practice in Golang
// even when that word doesn't exist
type Servicer interface {
	RegisterOnServer(*grpc.Server)
}

func Register(impl Servicer) {
	impl.RegisterOnServer(GRPCServer)
}
