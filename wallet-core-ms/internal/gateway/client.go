package gateway

import "github.com/raphaelmb/fullcycle-walletcore-ms/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
