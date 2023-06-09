package mysql

import (
	"context"

	"github.com/kumin/go-tpc/services/customer_service/entities"
	"github.com/kumin/go-tpc/services/customer_service/infras"
	"github.com/kumin/go-tpc/services/customer_service/repos"
	"gorm.io/gorm/clause"
)

var _ repos.OrderRepo = &OrderMysqlRepo{}

type OrderMysqlRepo struct {
	db *infras.MysqlConnector
}

func NewOrderMysqlRepo(
	db *infras.MysqlConnector,
) *OrderMysqlRepo {
	return &OrderMysqlRepo{
		db: db,
	}
}

func (o *OrderMysqlRepo) AddOrder(
	ctx context.Context,
	order *entities.Order,
) (*entities.Order, error) {
	if err := o.db.Client.WithContext(ctx).Create(order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (o *OrderMysqlRepo) UpdateOrderStatus(
	ctx context.Context,
	id int,
	status entities.OrderStatus,
) (*entities.Order, error) {
	var order entities.Order
	if err := o.db.Client.WithContext(ctx).
		Model(&order).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Update("status", int(status)).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
