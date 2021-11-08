package usecase

import (
	"accounts/internal/domain/entity"
	"accounts/internal/domain/repository"
	"context"
)

type (
	//GetAccountUseCase will fetch account(s) in the repository
	GetAccountUseCase struct {
		repositoryRegistry repository.Registry
	}
)

//NewGetAccountUseCase builds GetAccountUseCase with its dependencies
func NewGetAccountUseCase(repositoryRegistry repository.Registry) GetAccountUseCase {
	return GetAccountUseCase{repositoryRegistry: repositoryRegistry}
}

//ById fetches the account in the repository by its id
func (g GetAccountUseCase) ById(id int64) (entity.Account, error) {
	accRepository := g.repositoryRegistry.AccountReadOnlyRepository()

	return accRepository.FindById(context.Background(), id)
}
