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
	stmt, err := b.DB.Prepare("INSERT INTO balances (account_id, amount) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(balance.AccountID, balance.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (b *BalanceDB) GetByID(id string) (*entity.Balance, error) {
	balance := &entity.Balance{}
	stmt, err := b.DB.Prepare("SELECT * FROM balances WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&balance.ID, &balance.AccountID, &balance.Amount, &balance.CreatedAt, &balance.UpdatedAt); err != nil {
		return nil, err
	}
	return balance, nil
}
