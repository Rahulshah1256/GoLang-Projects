package api

import (
	"context"
	pb "poc_admission/pkg/pb"
)

func (server *Servers) GetRegInfo(ctx context.Context, req *pb.GetReg) (*pb.RegistraionRequest, error) {
	return &pb.RegistraionRequest{FullName: "Sandeep", Username: req.Username, Password: req.GetPassword(), Email: "123"}, nil
}

func (server *Servers) CreateRegistration(ctx context.Context, req *pb.RegistraionRequest) (*pb.RegistraionResponse, error) {
	return &pb.RegistraionResponse{Message: "done"}, nil
}
