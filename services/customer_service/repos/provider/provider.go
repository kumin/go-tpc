package provider

import (
	"github.com/google/wire"

	"github.com/kumin/go-tpc/services/customer_service/infras"
	"github.com/kumin/go-tpc/services/customer_service/repos/mysql"
)

var MysqlGraphSet = wire.NewSet(
	infras.InfaGraphSet,
	mysql.NewProductMysqlRepo,
)
