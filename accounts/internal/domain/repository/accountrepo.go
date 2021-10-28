package repository

import "accounts/internal/domain/entity"

// Account defines the contract to save/get information from Account domain
type Account interface {
	Save(account entity.Account) (entity.Account, error)
}
