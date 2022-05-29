package order

import (
	"context"

	"orderservice/model"
)

type IOrderRepository interface {
	GetOrder(ctx context.Context, orderID int64) (model.Order, error)
}

type OrderService struct {
	order IOrderRepository
}

// New will create the OrderService object.
// Params:
// @o = Order Repository
func New(o IOrderRepository) *OrderService {
	return &OrderService{
		order: o,
	}
}

func (o OrderService) GetOrder(ctx context.Context, orderID int64) (model.Order, error) {
	// Put business logic here if needed
	return o.order.GetOrder(ctx, orderID)
}
