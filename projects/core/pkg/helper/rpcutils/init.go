package rpcutils

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Get(ctx context.Context, field string) string {
	var value string
	fieldValues := GetAll(ctx, field)

	if len(fieldValues) > 0 {
		value = fieldValues[0]
	}

	return value
}

func GetAll(ctx context.Context, field string) []string {
	var value []string

	meta, ok := metadata.FromIncomingContext(ctx)
	if ok {
		value = meta[field]
	}

	return value
}

func SendAuthorizeError() error {
	return status.Error(codes.Unauthenticated, "not authorized")
}

func SendNotFoundError(msg string) error {
	return status.Error(codes.NotFound, msg)
}

func SendAbortError(err error) error {
	return status.Error(codes.Aborted, err.Error())
}
