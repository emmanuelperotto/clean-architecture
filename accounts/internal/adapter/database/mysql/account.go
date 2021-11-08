package mysql

import (
	"accounts/internal/adapter/database"
	"accounts/internal/domain/entity"
	"context"
	"log"
)

type (
	accountWriteOnlyRepository struct {
		db database.SQLExecutor
	}

	accountReadOnlyRepository struct {
		db database.SQLExecutor
	}
)

//newAccountWriteOnlyRepository builds write only repository with its dependencies
func newAccountWriteOnlyRepository(db database.SQLExecutor) accountWriteOnlyRepository {
	return accountWriteOnlyRepository{db: db}
}

//newAccountReadOnlyRepository builds read only repository with its dependencies
func newAccountReadOnlyRepository(db database.SQLExecutor) accountReadOnlyRepository {
	return accountReadOnlyRepository{db: db}
}

//Create executes INSERT operation into Accounts table
func (r accountWriteOnlyRepository) Create(ctx context.Context, account entity.Account) (entity.Account, error) {
	result, err := r.db.ExecContext(ctx, "INSERT INTO Accounts (DocumentNumber) VALUES (?)", account.DocumentNumber)
	if err != nil {
		return entity.Account{}, err
	}

	account.Id, _ = result.LastInsertId()

	return account, nil
}

//FindById performs a SELECT operation given an account ID
func (r accountReadOnlyRepository) FindById(ctx context.Context, id int64) (account entity.Account, err error) {
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
