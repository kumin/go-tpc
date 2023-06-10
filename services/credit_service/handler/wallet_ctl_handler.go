package handler

import (
	"net/http"

	"github.com/kumin/go-tpc/services/credit_service/services"
)

type WalletCtlHandler struct {
	walletService *services.WalletService
}

func NewWalletCtlHandler(
	walletService *services.WalletService,
) *WalletCtlHandler {
	return &WalletCtlHandler{
		walletService: walletService,
	}
}

func (w *WalletCtlHandler) AddWallet(req *http.Request) (interface{}, error) {
	return w.walletService.AddWallet(req)
}

func (w *WalletCtlHandler) UpdateWalletBalance(req *http.Request) (interface{}, error) {
	return w.walletService.UpdateWalletBalance(req)
}
