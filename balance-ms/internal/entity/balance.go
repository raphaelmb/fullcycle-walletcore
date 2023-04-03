package entity

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        string
	AccountID string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBalance(accountId string, amount float64) *Balance {
	return &Balance{
		ID:        uuid.New().String(),
		AccountID: accountId,
		Amount:    amount,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
