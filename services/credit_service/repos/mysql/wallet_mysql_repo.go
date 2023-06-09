package mysql

import (
	"context"
	"log"

	"github.com/kumin/go-tpc/services/credit_service/entities"
	"github.com/kumin/go-tpc/services/credit_service/repos"
	"github.com/kumin/go-tpc/services/customer_service/infras"
	"go.uber.org/zap"
)

var _ repos.WalletRepo = &WalletMysqlRepo{}

type WalletMysqlRepo struct {
	db     *infras.MysqlConnector
	logger *zap.Logger
}

func NewOrderMysqlRepo(
	db *infras.MysqlConnector,
) *WalletMysqlRepo {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	return &WalletMysqlRepo{
		db:     db,
		logger: logger,
	}
}

func (o *WalletMysqlRepo) AddWallet(
	ctx context.Context,
	wallet *entities.Wallet,
) (*entities.Wallet, error) {
	if err := o.db.Client.WithContext(ctx).Create(wallet).Error; err != nil {
		return nil, err
	}

	return wallet, nil
}

func (o *WalletMysqlRepo) UpdateWalletBalance(
	ctx context.Context,
	id int,
	amount float64,
) (*entities.Wallet, error) {
	var wallet entities.Wallet
	if err := o.db.Client.WithContext(ctx).
		Model(&wallet).
		Where("id = ?", id).
		Find(&wallet).Error; err != nil {
		return nil, err
	}

	if wallet.Money < amount {
		o.logger.Error(entities.InternalError.Error(), zap.Int("wallet id", id))
		return nil, entities.ErrNotEnoughBalance
	}
	wallet.Money -= amount
	if err := o.db.Client.WithContext(ctx).
		Model(&wallet).
		Select("money").
		Updates(wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}
