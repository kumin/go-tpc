package services

import "github.com/kumin/go-tpc/pkg/kafka"

type OrderCoordinatorService struct {
	producer *kafka.Producer
	consumer *kafka.Consumer
}

func NewOrderCoordinatorService() *OrderCoordinatorService {
	return &OrderCoordinatorService{}
}
