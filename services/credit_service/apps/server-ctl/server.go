package apps

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/kumin/go-tpc/services/credit_service/configs"
	"github.com/kumin/go-tpc/services/credit_service/handler"
)

type HttpServer struct {
	port   int
	logger *log.Logger
}

func NewHttpServer(
	configs *configs.ServerConfiguration,
	walletHandler *handler.WalletCtlHandler,
) *HttpServer {
	server := &HttpServer{
		port:   configs.Port,
		logger: log.Default(),
	}
	server.RegisterHandler("/v1/wallet/add", handler.HandlerWrapper(walletHandler.AddWallet))
	server.RegisterHandler("/v1/wallet/balance/update", handler.HandlerWrapper(walletHandler.UpdateWalletBalance))
	return server
}

func (h *HttpServer) Start(ctx context.Context) error {
	h.logger.Printf("Server is listening on port:%d", h.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", h.port), nil)
}

func (h *HttpServer) RegisterHandler(path string, handler http.HandlerFunc) {
	http.Handle(path, handler)
}
