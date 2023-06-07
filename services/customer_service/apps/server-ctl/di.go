package apps

import (
	"github.com/google/wire"
	"github.com/kumin/go-tpc/services/customer_service/configs"
)

var ServerGraphSet = wire.NewSet(
	configs.ConfigGraphSet,
	NewHttpServer,
)
