package usecase

import (
	"accounts/internal/domain/entity"
	"errors"
	"math/rand"
)

//CreateAccountUseCase receives a creation request and return a response
type CreateAccountUseCase struct {
}

//CreateAccountRequest is the request object to be used in the CreateAccountUseCase
type CreateAccountRequest struct {
	DocumentNumber string
}

//Call performs the account creation use case
func (c CreateAccountUseCase) Call(request CreateAccountRequest) (entity.Account, error) {
	if request.DocumentNumber == "" {
		return entity.Account{}, errors.New("DocumentNumber required")
	}

	account := entity.Account{
		Id:             rand.Int63(),
		DocumentNumber: request.DocumentNumber,
	}

	return account, nil
}
