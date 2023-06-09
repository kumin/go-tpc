package repos

import (
	"context"

	"github.com/kumin/go-tpc/services/credit_service/entities"
)

type WalletRepo interface {
	AddWallet(ctx context.Context, wallet *entities.Wallet) (*entities.Wallet, error)
	UpdateWalletBalance(ctx context.Context, id int, amount float64) (*entities.Wallet, error)
}
