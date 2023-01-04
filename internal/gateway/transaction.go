package gateway

import "github.com/raphaelmb/fullcycle-walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
