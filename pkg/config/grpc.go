package config

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection (addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr, 
		grpc.WithTimeout(5 * time.Second),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Printf("Err connecting grpc server at %s, %s", addr, err)	
	}
	return conn
}
