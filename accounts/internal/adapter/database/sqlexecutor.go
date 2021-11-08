package database

import (
	"context"
	"database/sql"
)

//SQLExecutor is an interface that makes possible to sql.DB and sql.Tx implement it
type SQLExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
