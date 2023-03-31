package mocks

import (
	"context"

	"github.com/raphaelmb/fullcycle-walletcore-ms/pkg/uow"
	"github.com/stretchr/testify/mock"
)

type UowMock struct {
	mock.Mock
}

func (m *UowMock) Register(name string, fc uow.RepositoryFactory) {
	m.Called(name, fc)
}

func (m *UowMock) GetRepository(ctx context.Context, name string) (any, error) {
	args := m.Called(ctx, name)
	return args.Get(0), args.Error(1)
}

func (m *UowMock) Do(ctx context.Context, fn func(uow *uow.Uow) error) error {
	args := m.Called(fn)
	return args.Error(0)
}

func (m *UowMock) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *UowMock) UnRegister(name string) {
	m.Called(name)
}

func (m *UowMock) CommitOrRollback() error {
	return nil
}
