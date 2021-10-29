package usecase

import (
	"accounts/internal/domain/entity"
	"accounts/internal/domain/repository"
)

type (
	//GetAccountUseCase will fetch account(s) in the repository
	GetAccountUseCase struct {
		accountRepository repository.AccountReadOnly
	}
)

//NewGetAccountUseCase builds GetAccountUseCase with its dependencies
func NewGetAccountUseCase(accountRepository repository.AccountReadOnly) GetAccountUseCase {
	return GetAccountUseCase{accountRepository: accountRepository}
}

//ById fetches the account in the repository by its id
func (g GetAccountUseCase) ById(id int64) (entity.Account, error) {
	return g.accountRepository.FindById(id)
}
