package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

//ConnectMySQLDB starts a connection using MySQL driver. Returns an error if it can't establish the connection
func ConnectMySQLDB() (db *sql.DB, err error) {
	const (
		driver   = "mysql"
		user     = "example"
		password = "secret123"
		host     = "localhost"
		port     = "3306"
		dbName   = "accounts"
	)

	db, err = sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName))

	if err != nil {
		log.Println("Error opening DB", err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error pinging DB", err.Error())
		return
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(50)

	return
}
