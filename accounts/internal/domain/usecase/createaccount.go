package usecase

import (
	"accounts/internal/domain/entity"
	"accounts/internal/domain/repository"
	"context"
	"errors"
)

type (
	//CreateAccountUseCase receives a creation request and return a response
	CreateAccountUseCase struct {
		repositoryRegistry repository.Registry
	}

	//CreateAccountRequest is the request object to be used in the CreateAccountUseCase
	CreateAccountRequest struct {
		DocumentNumber string
	}
)

//NewCreateAccountUseCase builds CreateAccountUseCase with its dependencies
func NewCreateAccountUseCase(registry repository.Registry) CreateAccountUseCase {
	return CreateAccountUseCase{repositoryRegistry: registry}
}

//Call performs the account creation use case
func (c CreateAccountUseCase) Call(ctx context.Context, request CreateAccountRequest) (entity.Account, error) {
	if request.DocumentNumber == "" {
		return entity.Account{}, errors.New("DocumentNumber required")
	}

	account, err := c.repositoryRegistry.DoInTransaction(func(registry repository.Registry) (interface{}, error) {
		accRepository := c.repositoryRegistry.AccountWriteOnlyRepository()
		account := entity.Account{
			DocumentNumber: request.DocumentNumber,
		}

		return accRepository.Create(ctx, account)
	})

	if err != nil {
		return entity.Account{}, err
	}

	return account.(entity.Account), nil
}
