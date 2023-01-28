package config

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection () *grpc.ClientConn {
	conn, err := grpc.Dial(
		"localhost:8081", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		panic(err)
	}
	return conn
}
