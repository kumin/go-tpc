package entities

type Wallet struct {
	Id     int     `json:"id,omitempty"`
	UserId int     `json:"user_id,omitempty"`
	Money  float64 `json:"money,omitempty"`
}

func (w *Wallet) TableName() string {
	return "wallet"
}
