package handler

import (
	"context"
	"microservices-gRPC-go/services/common/genproto/orders"
	"microservices-gRPC-go/services/orders/types"
)

type OrderGrpcHandler struct {
	//service injection
	ordersService types.OrderService
	// uniplmeneted UnimplementedOrderServiceServer
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService() {
	gRPCHandler := &OrderGrpcHandler{}

	// Register the OrderServiceServer
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.ordersService.CreateOrder(ctx, order)

	if err != nil {
		return nil, err
	}
	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}
