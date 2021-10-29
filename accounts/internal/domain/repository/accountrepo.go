package repository

import "accounts/internal/domain/entity"

// AccountReadOnly defines the contract to get information from Account domain
type AccountReadOnly interface {
	FindById(id int64) (entity.Account, error)
}

// AccountWriteOnly defines the contract to save/update information from Account domain
type AccountWriteOnly interface {
	Save(account entity.Account) (entity.Account, error)
}
