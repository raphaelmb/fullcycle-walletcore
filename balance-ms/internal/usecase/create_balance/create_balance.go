package createbalance

import (
	"github.com/raphaelmb/fullcycle-balance-ms/internal/entity"
	"github.com/raphaelmb/fullcycle-balance-ms/internal/gateway"
)

type CreateBalanceInputDTO struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type CreateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewCreateBalanceUseCase(b gateway.BalanceGateway) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{
		BalanceGateway: b,
	}
}

func (uc *CreateBalanceUseCase) Execute(input CreateBalanceInputDTO) error {
	balance := entity.NewBalance(input.AccountID, input.Amount)
	err := uc.BalanceGateway.Save(balance)
	if err != nil {
		return err
	}
	return nil
}
