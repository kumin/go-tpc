//go:build wireinject
// +build wireinject

package apps

import (
	"github.com/google/wire"
	"github.com/kumin/go-tpc/services/customer_service/handler"
	"github.com/kumin/go-tpc/services/customer_service/repos/provider"
	"github.com/kumin/go-tpc/services/customer_service/services"
)

var SuperGraphSet = wire.NewSet(
	provider.MysqlGraphSet,
	services.ServiceGraphSet,
	handler.HandlerGraphSet,
	ServerGraphSet,
)

func BuildServer() (*HttpServer, error) {
	wire.Build(
		SuperGraphSet,
	)

	return nil, nil
}
