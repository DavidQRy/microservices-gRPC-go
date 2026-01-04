package main

import (
	"log"
	handler "microservices-gRPC-go/services/orders/handler/orders"
	"microservices-gRPC-go/services/orders/service"
	"net"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Println("Starting gRPC server on ", s.addr)

	return grpcServer.Serve(lis)
}
