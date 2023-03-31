package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestNewAccountWithNilClient(t *testing.T) {
	client, _ := NewClient("", "")
	account := NewAccount(client)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.NotNil(t, account)
	assert.Equal(t, account.Balance, 100.0)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)
	assert.NotNil(t, account)
	assert.Equal(t, account.Balance, 50.0)
}
