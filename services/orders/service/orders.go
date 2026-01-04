package service

import (
	"context"
	"microservices-gRPC-go/services/common/genproto/orders"
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
