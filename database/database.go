package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbCredentialsFormat = "user=%s password=%s dbname=%s host=%s port=%d"
)

func GetDatabaseConnection(cfg config) *sql.DB {
	address := fmt.Sprintf(dbCredentialsFormat,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Host,
		cfg.Port,
	)

	db, err := sql.Open("mysql", address)
	if err != nil {
		log.Fatal("[Database] failed connecting to DB: " + address + ", err: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal("[Database] db is unreachable: " + address + ", err: " + err.Error())
	}

	return db
}
