package server

import (
	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	server := grpc.NewServer()

	return server
}