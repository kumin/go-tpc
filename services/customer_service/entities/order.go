package entities

type OrderStatus int

const (
	PeddingStatus    OrderStatus = 1
	SuccessfulStatus OrderStatus = 2
	FailedStatus     OrderStatus = 3
)

type Order struct {
	Id              int         `json:"id,omitempty"`
	Products        []*Product  `json:"products,omitempty"`
	Money           float64     `json:"money,omitempty"`
	ShippingAddress string      `json:"shipping_address,omitempty"`
	Status          OrderStatus `json:"status,omitempty"`
}
