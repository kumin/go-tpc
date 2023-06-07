package services

import (
	"context"

	"github.com/kumin/go-tpc/services/customer_service/entities"
)

var prods = []*entities.Product{
	{
		ID:    123456,
		Title: "Iphone 13",
		Properties: &entities.Properties{
			Price:    30000000.04,
			Category: "Smart Phone",
		},
	},
	{
		ID:    123458,
		Title: "Tivi Sony",
		Properties: &entities.Properties{
			Price:    15000000.04,
			Category: "Tivi",
		},
	},
	{
		ID:    123457,
		Title: "Tu Lanh Panasonic",
		Properties: &entities.Properties{
			Price:    20000000.08,
			Category: "Tu Lanh",
		},
	},
}

type ProductCtlMockService struct{}

func NewProductCtlHandler() *ProductCtlMockService {
	return &ProductCtlMockService{}
}

func (p *ProductCtlMockService) ListProducts(
	ctx context.Context,
	page int,
	limit int) (
	[]*entities.Product, error,
) {
	return prods, nil
}

func (p *ProductCtlMockService) GetProduct(
	ctx context.Context,
	id int64,
) (*entities.Product, error) {
	for _, p := range prods {
		if p.ID == id {
			return p, nil
		}
	}

	return nil, nil
}
