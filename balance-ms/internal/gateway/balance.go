package gateway

import "github.com/raphaelmb/fullcycle-balance-ms/internal/entity"

type BalanceGateway interface {
	List(id string) (*entity.Balance, error)
	Save(balance *entity.Balance) error
}
