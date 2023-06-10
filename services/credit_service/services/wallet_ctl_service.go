package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/kumin/go-tpc/services/credit_service/entities"
	"github.com/kumin/go-tpc/services/credit_service/repos"
)

type UpdateBody struct {
	Id     int     `json:"id,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type WalletService struct {
	walletRepo repos.WalletRepo
}

func NewWalletService(
	walletRepo repos.WalletRepo,
) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
	}
}

func (o *WalletService) AddWallet(req *http.Request) (*entities.Wallet, error) {
	d, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var wallet entities.Wallet
	if err := json.Unmarshal(d, &wallet); err != nil {
		return nil, err
	}
	return o.walletRepo.AddWallet(req.Context(), &wallet)
}

func (o *WalletService) UpdateWalletBalance(req *http.Request) (*entities.Wallet, error) {
	d, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var updateBody UpdateBody
	if err := json.Unmarshal(d, &updateBody); err != nil {
		return nil, err
	}
	if updateBody.Id == 0 || updateBody.Amount == 0 {
		return nil, entities.ParamInvalid
	}
	return o.walletRepo.UpdateWalletBalance(req.Context(),
		updateBody.Id, updateBody.Amount)
}
