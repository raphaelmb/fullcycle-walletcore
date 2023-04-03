package createbalance

import (
	"github.com/raphaelmb/fullcycle-balance-ms/internal/entity"
	"github.com/raphaelmb/fullcycle-balance-ms/internal/gateway"
)

type CreateBalanceInputDTO struct {
	Balance *entity.Balance
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
	err := uc.BalanceGateway.Save(input.Balance)
	if err != nil {
		return err
	}
	return nil
}
