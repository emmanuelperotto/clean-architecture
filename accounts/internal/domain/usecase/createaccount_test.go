package usecase_test

import (
	"accounts/internal/domain/entity"
	"accounts/internal/domain/repository"
	"accounts/internal/domain/usecase"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type repoRegistryMock struct {
	mock.Mock
}

type accReadOnlyMock struct {
	mock.Mock
}

type accWriteOnlyMock struct {
	mock.Mock
}

func (a *accReadOnlyMock) FindById(ctx context.Context, id string) (entity.Account, error) {
	args := a.Called(ctx, id)
	return args.Get(0).(entity.Account), args.Error(1)
}

func (a *accWriteOnlyMock) Create(ctx context.Context, account entity.Account) (entity.Account, error) {
	args := a.Called(ctx, account)
	return args.Get(0).(entity.Account), args.Error(1)
}

func (r *repoRegistryMock) AccountReadOnlyRepository() repository.AccountReadOnly {
	args := r.Called()
	return args.Get(0).(repository.AccountReadOnly)
}

func (r *repoRegistryMock) AccountWriteOnlyRepository() repository.AccountWriteOnly {
	args := r.Called()
	return args.Get(0).(repository.AccountWriteOnly)
}

func (r *repoRegistryMock) DoInTransaction(txFunc repository.InTransaction) (out interface{}, err error) {
	args := r.Called(txFunc)

	return args.Get(0).(repository.InTransaction)(r)
}

func TestNewCreateAccountUseCase(t *testing.T) {
	repoRegistry := repoRegistryMock{}
	accReadOnly := accReadOnlyMock{}
	accWriteOnly := accWriteOnlyMock{}

	repoRegistry.On("AccountReadOnlyRepository").Return(&accReadOnly)
	repoRegistry.On("AccountWriteOnlyRepository").Return(&accWriteOnly)

	useCase := usecase.NewCreateAccountUseCase(&repoRegistry)

	assert.NotNil(t, useCase)
}
