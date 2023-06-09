package repos

import (
	"context"

	"github.com/kumin/go-tpc/services/customer_service/entities"
)

type OrderRepo interface {
	AddOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	UpdateOrderStatus(ctx context.Context, id int, status entities.OrderStatus) (*entities.Order, error)
}
