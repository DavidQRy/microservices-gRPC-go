package service

import (
	"context"
	"log"
	"microservices-gRPC-go/services/common/genproto/orders"
	handler "microservices-gRPC-go/services/orders/handler/orders"
	"microservices-gRPC-go/services/orders/service"
	"net"

	"google.golang.org/grpc"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct {
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDb = append(ordersDb, order)

	return nil
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
