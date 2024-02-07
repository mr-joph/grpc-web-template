package register

import (
	"context"

	"core/pkg/proto/pb"
)

type LiveclassServer struct {
	pb.UnimplementedLiveclassServiceServer
}

func (s *LiveclassServer) List(
	ctx context.Context,
	_ *pb.LiveclassIn,
) (*pb.LiveclassOut, error) {
	return &pb.LiveclassOut{Name: "test"}, nil
}

func init() {
	pb.RegisterLiveclassServiceServer(GRPCServer, &LiveclassServer{})
}
