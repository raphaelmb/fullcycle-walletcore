package database

import (
	"database/sql"

	"github.com/raphaelmb/fullcycle-balance-ms/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (b *BalanceDB) Save(balance *entity.Balance) error {
	return nil
}

func (b *BalanceDB) List(id string) (*entity.Balance, error) {
	var balance entity.Balance
	stmt, err := b.DB.Prepare("SELECT * FROM balances WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&balance.ID, &balance.AccountID, &balance.Amount, &balance.CreatedAt, &balance.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}
