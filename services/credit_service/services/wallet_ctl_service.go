package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kumin/go-tpc/services/customer_service/entities"
	"github.com/kumin/go-tpc/services/customer_service/repos"
)

type UpdateBody struct {
	Id     int     `json:"id,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type OrderService struct {
	orderRepo repos.OrderRepo
}

func NewOrderService(
	orderRepo repos.OrderRepo,
) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (o *OrderService) AddOrder(req *http.Request) (*entities.Order, error) {
	d, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var order entities.Order
	if err := json.Unmarshal(d, &order); err != nil {
		return nil, err
	}
	return o.orderRepo.AddOrder(req.Context(), &order)
}

func (o *OrderService) UpdateOrderStatus(req *http.Request) (*entities.Order, error) {
	d, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var updateBody UpdateBody
	if err := json.Unmarshal(d, &updateBody); err != nil {
		return nil, err
	}
	if updateBody.Id == 0 || updateBody.Status == 0 {
		return nil, entities.ParamInvalid
	}
	return o.orderRepo.UpdateOrderStatus(req.Context(),
		updateBody.Id, entities.OrderStatus(updateBody.Status))
}
