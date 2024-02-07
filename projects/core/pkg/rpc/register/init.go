package register

import (
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LogIntercept(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (res interface{}, err error) {

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not grab metadata from context")
	}

	fmt.Println("Log inteceptor")
	log.Printf("%v", info.FullMethod)
	fmt.Println(req)
	fmt.Println(meta)

	return handler(ctx, req)
}

func AuthIntercept(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (res interface{}, err error) {

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("could not grab metadata from context")
	}
	fmt.Println("Auth inteceptor")
	log.Printf("%v", info.FullMethod)
	fmt.Println(req)
	fmt.Println(meta)

	return handler(ctx, req)
}

var GRPCServer = grpc.NewServer(
	grpc.ChainUnaryInterceptor(
		LogIntercept,
		AuthIntercept,
	),
)
