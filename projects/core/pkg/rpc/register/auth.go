package register

import (
	"context"

	"core/pkg/proto/pb"
)

type CredentialServer struct {
	pb.UnimplementedCredentialServiceServer
}

func (s *CredentialServer) Login(
	ctx context.Context,
	req *pb.CredentialRequest,
) (*pb.CredentialResponse, error) {
	return nil, nil
}

func init() {
	pb.RegisterCredentialServiceServer(GRPCServer, &CredentialServer{})
}
