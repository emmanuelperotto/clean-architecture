package local

import (
	"accounts/internal/domain/entity"
	"context"
	"errors"
	"log"
	"math/rand"
)

//accountStore stores in memory
var accountStore []entity.Account

type (
	// accountWriteOnlyRepository represents in memory write only repository for Account
	accountWriteOnlyRepository struct{}

	// accountReadOnlyRepository represents in memory read only repository for Account
	accountReadOnlyRepository struct{}
)

//Create persists the account entity in memory
func (r accountWriteOnlyRepository) Create(_ context.Context, account entity.Account) (entity.Account, error) {
	account.Id = rand.Int63()

	accountStore = append(accountStore, account)

	log.Println("Account created: ", account.Id)
	log.Println("Current store len: ", len(accountStore))

	return account, nil
}

//FindById fetches the account entity in memory given an id
func (r accountReadOnlyRepository) FindById(_ context.Context, id int64) (entity.Account, error) {
	for _, account := range accountStore {
		if account.Id == id {
			return account, nil
		}
	}

	return entity.Account{}, errors.New("account not present")
}
