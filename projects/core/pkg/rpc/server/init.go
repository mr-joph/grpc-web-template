package server

import "google.golang.org/grpc"

var GRPCServer = grpc.NewServer(
	grpc.ChainUnaryInterceptor(
		LogIntercept,
		AuthIntercept,
	),
)
