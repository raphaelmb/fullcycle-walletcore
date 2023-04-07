package getbalance

import (
	"time"

	"github.com/raphaelmb/fullcycle-balance-ms/internal/gateway"
)

type GetBalanceInputDTO struct {
	ID string
}

type GetBalanceOutputDTO struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetBalanceByIDUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceByIDUseCase(b gateway.BalanceGateway) *GetBalanceByIDUseCase {
	return &GetBalanceByIDUseCase{
		BalanceGateway: b,
	}
}

func (uc *GetBalanceByIDUseCase) Execute(input GetBalanceInputDTO) (*GetBalanceOutputDTO, error) {
	balance, err := uc.BalanceGateway.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetBalanceOutputDTO{
		ID:        balance.ID,
		AccountID: balance.AccountID,
		Amount:    balance.Amount,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}, nil
}
