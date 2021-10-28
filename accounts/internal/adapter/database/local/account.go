package local

import (
	"accounts/internal/domain/entity"
	"log"
	"math/rand"
)

//accountStore stores in memory
var accountStore []entity.Account

// AccountRepository represents in memory repository for Account
type AccountRepository struct {
}

//NewAccountRepository builds the repository with its dependencies
func NewAccountRepository() AccountRepository {
	return AccountRepository{}
}

//Save persists the account entity in memory
func (r AccountRepository) Save(account entity.Account) (entity.Account, error) {
	account.Id = rand.Int63()

	accountStore = append(accountStore, account)

	log.Println("Account created: ", account.Id)
	log.Println("Current store len: ", len(accountStore))

	return account, nil
}
