package apps

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/kumin/go-tpc/services/customer_service/configs"
	"github.com/kumin/go-tpc/services/customer_service/handler"
)

type HttpServer struct {
	port   int
	logger *log.Logger
}

func NewHttpServer(
	configs *configs.ServerConfiguration,
	prodHandler *handler.ProductCtlHandler,
	orderHandler *handler.OrderCtlHandler,
) *HttpServer {
	server := &HttpServer{
		port:   configs.Port,
		logger: log.Default(),
	}
	server.RegisterHandler("/v1/product/add", handler.HandlerWrapper(prodHandler.AddProduct))
	server.RegisterHandler("/v1/product/listing", handler.HandlerWrapper(prodHandler.ListProducts))
	server.RegisterHandler("/v1/product", handler.HandlerWrapper(prodHandler.GetProduct))
	server.RegisterHandler("/v1/order", handler.HandlerWrapper(orderHandler.AddOrder))
	server.RegisterHandler("/v1/order/status/update", handler.HandlerWrapper(orderHandler.UpdateOrderStatus))
	return server
}

func (h *HttpServer) Start(ctx context.Context) error {
	h.logger.Printf("Server is listening on port:%d", h.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", h.port), nil)
}

func (h *HttpServer) RegisterHandler(path string, handler http.HandlerFunc) {
	http.Handle(path, handler)
}
