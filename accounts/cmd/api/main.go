package main

import (
	"accounts/internal/adapter/database/local"
	"accounts/internal/adapter/web"
	"accounts/internal/domain/usecase"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatalln("Fatal Error", err.Error())
	}

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			log.Println("Can't close DB Connection")
		}
	}(db)

	app := fiber.New()
	accountWriteOnlyRepository := local.NewAccountWriteOnlyRepository()
	accountReadOnlyRepository := local.NewAccountReadOnlyRepository()

	accountHandler := web.NewAccountHandler(
		usecase.NewCreateAccountUseCase(accountWriteOnlyRepository),
		usecase.NewGetAccountUseCase(accountReadOnlyRepository),
	)

	app.Post("/accounts", accountHandler.CreateAccount)
	app.Get("/accounts/:id", accountHandler.GetAccount)

	err = app.Listen(":3000")
	if err != nil {
		log.Fatalln("Error starting server: ", err.Error())
	}
}

func setupDB() (db *sql.DB, err error) {
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
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return
}
