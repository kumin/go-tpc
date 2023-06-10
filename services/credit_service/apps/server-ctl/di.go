package apps

import (
	"github.com/google/wire"
	"github.com/kumin/go-tpc/services/credit_service/configs"
)

var ServerGraphSet = wire.NewSet(
	configs.ConfigGraphSet,
	NewHttpServer,
)
