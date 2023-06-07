package repos

import (
	"context"
)

type OrderRepo interface {
	AddOrder(ctx context.Context, order *Order) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) (*Order, error)
}
