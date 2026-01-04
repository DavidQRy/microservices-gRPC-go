package types

import (
	"context"
	"microservices-gRPC-go/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
