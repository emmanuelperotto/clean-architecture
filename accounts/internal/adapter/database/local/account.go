package local

import (
	"accounts/internal/domain/entity"
	"errors"
	"log"
	"math/rand"
)

//accountStore stores in memory
var accountStore []entity.Account

type (
	// AccountWriteOnlyRepository represents in memory write only repository for Account
	AccountWriteOnlyRepository struct{}

	// AccountReadOnlyRepository represents in memory read only repository for Account
	AccountReadOnlyRepository struct{}
)

//NewAccountWriteOnlyRepository builds write only repository with its dependencies
func NewAccountWriteOnlyRepository() AccountWriteOnlyRepository {
	return AccountWriteOnlyRepository{}
}

//NewAccountReadOnlyRepository builds read only repository with its dependencies
func NewAccountReadOnlyRepository() AccountReadOnlyRepository {
	return AccountReadOnlyRepository{}
}

//Save persists the account entity in memory
func (r AccountWriteOnlyRepository) Save(account entity.Account) (entity.Account, error) {
	account.Id = rand.Int63()

	accountStore = append(accountStore, account)

	log.Println("Account created: ", account.Id)
	log.Println("Current store len: ", len(accountStore))

	return account, nil
}

//FindById fetches the account entity in memory given an id
func (a AccountReadOnlyRepository) FindById(id int64) (entity.Account, error) {
	for _, account := range accountStore {
		if account.Id == id {
			return account, nil
		}
	}

	return entity.Account{}, errors.New("account not present")

}
