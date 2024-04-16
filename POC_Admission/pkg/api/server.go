package api

import (
	"fmt"
	pb "poc_admission/pkg/pb"
	"poc_admission/pkg/token"
	"poc_admission/pkg/util"
)

type Servers struct {
	pb.UnimplementedAdmissionServiceServer
	config     util.Config
	tokenMaker token.Maker
}

func newServer(config util.Config) (*Servers, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// server := &Server{
	// 	config:     util.Config,
	// 	tokenMaker: tokenMaker,
	// }

	return &Servers{}, nil
}
