package register

import (
	"context"
	"fmt"

	"core/pkg/proto/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) List(
	ctx context.Context,
	req *pb.UserIn,
) (*pb.UserOut, error) {
	return &pb.UserOut{Name: req.Name}, nil
}

func (s *UserServer) Filter(
	ctx context.Context,
	_ *empty.Empty,
) (*pb.UserOut, error) {
	return &pb.UserOut{Name: "Test"}, nil
}

func (s *UserServer) Auth(
	ctx context.Context,
	_ *pb.UserAuthIn,
) (*pb.UserAuthOut, error) {
	fmt.Println("Auth processing...")
	return &pb.UserAuthOut{
		Token: "1234567890",
	}, nil
}

func init() {
	pb.RegisterUserServiceServer(GRPCServer, &UserServer{})
}
