package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	if err != nil {
		t.Errorf("Error creating new client: %v", err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("John Doe Update", "j@g.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@g.com", client.Email)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("", "j@g.com")
	assert.Error(t, err, "name is required")
}

func TestAddAcount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	err := client.AddAcount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
	assert.Equal(t, client.Accounts[0], account)
}
