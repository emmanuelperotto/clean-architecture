package usecase

import (
	"accounts/internal/domain/entity"
	"accounts/internal/domain/repository"
	"errors"
)

type (
	//CreateAccountUseCase receives a creation request and return a response
	CreateAccountUseCase struct {
		accountRepository repository.Account
	}

	//CreateAccountRequest is the request object to be used in the CreateAccountUseCase
	CreateAccountRequest struct {
		DocumentNumber string
	}
)

//NewCreateAccountUseCase builds CreateAccountUseCase with its dependencies
func NewCreateAccountUseCase(accountRepository repository.Account) CreateAccountUseCase {
	return CreateAccountUseCase{accountRepository: accountRepository}
}

//Call performs the account creation use case
func (c CreateAccountUseCase) Call(request CreateAccountRequest) (entity.Account, error) {
	if request.DocumentNumber == "" {
		return entity.Account{}, errors.New("DocumentNumber required")
	}

	account := entity.Account{
		DocumentNumber: request.DocumentNumber,
	}

	return c.accountRepository.Save(account)
}
