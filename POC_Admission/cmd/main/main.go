package main

import (
	"net"
	pbserver "poc_admission/pkg/api"
	pb "poc_admission/pkg/pb"
	c "poc_admission/pkg/util"
	logger "poc_admission/pkg/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Resitration struct {
	pb.UnimplementedAdmissionServiceServer
}

func main() {
	configs, err := c.LoadConfig("../../")
	if err != nil {
		logger.Logger().Error("Cannot load config file")
	}
	lis, err := net.Listen("tcp", configs.Port)
	if err != nil {
		logger.Logger().Error(err.Error())
	}
	s := grpc.NewServer()

	pb.RegisterAdmissionServiceServer(s, &pbserver.Servers{})
	reflection.Register(s)
	logger.Logger().Info("server Started at " + lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		logger.Logger().Error(err.Error())
	}
}
