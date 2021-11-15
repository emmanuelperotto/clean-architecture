package repository

import (
	"accounts/internal/domain/entity"
	"context"
)

// AccountReadOnly defines the contract to get information from Account domain
type AccountReadOnly interface {
	FindById(ctx context.Context, id string) (entity.Account, error)
}

// AccountWriteOnly defines the contract to save/update information from Account domain
type AccountWriteOnly interface {
	Create(ctx context.Context, account entity.Account) (entity.Account, error)
}
