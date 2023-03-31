package gateway

import "github.com/raphaelmb/fullcycle-walletcore-ms/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
