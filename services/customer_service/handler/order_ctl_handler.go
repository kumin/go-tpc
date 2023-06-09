package handler

import (
	"net/http"

	"github.com/kumin/go-tpc/services/customer_service/services"
)

type OrderCtlHandler struct {
	orderService *services.OrderService
}

func NewOrderCtlHandler(
	orderService *services.OrderService,
) *OrderCtlHandler {
	return &OrderCtlHandler{
		orderService: orderService,
	}
}

func (o *OrderCtlHandler) AddOrder(
	req *http.Request,
) (interface{}, error) {
	return o.orderService.AddOrder(req)
}

func (o *OrderCtlHandler) UpdateOrderStatus(
	req *http.Request,
) (interface{}, error) {
	return o.orderService.UpdateOrderStatus(req)
}
