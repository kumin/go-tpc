package provider

import (
	"github.com/google/wire"

	"github.com/kumin/go-tpc/services/credit_service/repos"
	"github.com/kumin/go-tpc/services/credit_service/repos/mysql"
)

var MysqlGraphSet = wire.NewSet(
	mysql.NewWalletMysqlRepo,
	wire.Bind(new(repos.WalletRepo), new(*mysql.WalletMysqlRepo)),
)
