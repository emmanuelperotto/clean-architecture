package mysql

import (
	"accounts/internal/domain/entity"
	"context"
	"database/sql"
	"log"
)

type (
	// AccountWriteOnlyRepository represents in memory write only repository for Account
	AccountWriteOnlyRepository struct {
		db *sql.DB
	}

	// AccountReadOnlyRepository represents in memory read only repository for Account
	AccountReadOnlyRepository struct {
		db *sql.DB
	}
)

//NewAccountWriteOnlyRepository builds write only repository with its dependencies
func NewAccountWriteOnlyRepository(db *sql.DB) AccountWriteOnlyRepository {
	return AccountWriteOnlyRepository{db: db}
}

//NewAccountReadOnlyRepository builds read only repository with its dependencies
func NewAccountReadOnlyRepository(db *sql.DB) AccountReadOnlyRepository {
	return AccountReadOnlyRepository{db: db}
}

//Save persists the account entity in memory
func (r AccountWriteOnlyRepository) Save(ctx context.Context, account entity.Account) (entity.Account, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return entity.Account{}, err
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	result, err := tx.ExecContext(ctx, "INSERT INTO Accounts (DocumentNumber) VALUES (?)", account.DocumentNumber)
	if err != nil {
		return entity.Account{}, err
	}

	account.Id, _ = result.LastInsertId()

	return account, tx.Commit()
}

//FindById fetches the account entity in memory given an id
func (r AccountReadOnlyRepository) FindById(ctx context.Context, id int64) (account entity.Account, err error) {
	row := r.db.QueryRowContext(ctx, "SELECT Account_ID, DocumentNumber FROM Accounts a WHERE a.Account_ID = ?", id)

	var docNumber []byte

	if err = row.Scan(&id, &docNumber); err != nil {
		log.Println("Error finding Account by id", err.Error())
		return
	}

	account = entity.Account{
		Id:             id,
		DocumentNumber: string(docNumber),
	}

	return
}
