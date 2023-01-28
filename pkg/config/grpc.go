package config

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection (addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr, 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		panic(err)
	}
	return conn
}
