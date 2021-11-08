package mysql

import (
	"accounts/internal/adapter/database"
	"accounts/internal/domain/repository"
	"database/sql"
	"errors"
)

type mySQLRepositoryRegistry struct {
	db         *sql.DB
	dbExecutor database.SQLExecutor
}

//NewMySQLRepositoryRegistry builds mySQLRepositoryRegistry with its dependencies
func NewMySQLRepositoryRegistry(db *sql.DB) repository.Registry {
	return mySQLRepositoryRegistry{
		db: db,
	}
}

//GetAccountReadOnlyRepository returns the repository with its proper dbExecutor
func (r mySQLRepositoryRegistry) AccountReadOnlyRepository() repository.AccountReadOnly {
	if r.dbExecutor != nil {
		return newAccountReadOnlyRepository(r.dbExecutor)
	}

	return newAccountReadOnlyRepository(r.db)
}

//GetAccountWriteOnlyRepository returns the repository with its proper dbExecutor
func (r mySQLRepositoryRegistry) AccountWriteOnlyRepository() repository.AccountWriteOnly {
	if r.dbExecutor != nil {
		return newAccountWriteOnlyRepository(r.dbExecutor)
	}

	return newAccountWriteOnlyRepository(r.db)
}

//DoInTransaction Performs the operation in the txFunc using transactions, e.g. rollbacks everything if something fails
func (r mySQLRepositoryRegistry) DoInTransaction(txFunc repository.InTransaction) (out interface{}, err error) {
	var tx *sql.Tx
	registry := r
	if r.dbExecutor == nil {
		tx, err = r.db.Begin()
		if err != nil {
			return
		}
		defer func() {
			if p := recover(); p != nil {
				_ = tx.Rollback()
				panic(p) // re-throw panic after Rollback
			} else if err != nil {
				xerr := tx.Rollback() // err is non-nil; don't change it
				if xerr != nil {
					err = errors.New(xerr.Error())
				}
			} else {
				err = tx.Commit() // err is nil; if Commit returns error update err
			}
		}()
		registry = mySQLRepositoryRegistry{
			db:         r.db,
			dbExecutor: tx,
		}
	}
	out, err = txFunc(registry)
	return
}
