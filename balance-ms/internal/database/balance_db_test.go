package database

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/raphaelmb/fullcycle-balance-ms/internal/entity"
	"github.com/stretchr/testify/suite"
)

type BalanceDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	balanceDB *BalanceDB
}

func (s *BalanceDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE balances (id varchar(255), account_id varchar(255), amount float, created_at date, updated_at date)")
	s.balanceDB = NewBalanceDB(db)
}

func (s *BalanceDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE balances")
}

func TestBalanceDBTestSuite(t *testing.T) {
	suite.Run(t, new(BalanceDBTestSuite))
}

func (s *BalanceDBTestSuite) TestSave() {
	balance := &entity.Balance{
		ID:        "1",
		AccountID: "1",
		Amount:    10.0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.balanceDB.Save(balance)
	s.Nil(err)
}

// func (s *BalanceDBTestSuite) TestList() {
// 	balance := entity.NewBalance("1", 10.0)
// 	s.balanceDB.Save(balance)

// 	balanceDB, err := s.balanceDB.List(balance.ID)
// 	s.Nil(err)
// 	s.Equal(balance.ID, balanceDB.ID)
// 	s.Equal(balance.Amount, balanceDB.Amount)
// }
