package repository

import (
	"database/sql"
	"fmt"
	"log"
	"service-transaction/src"
	"service-transaction/src/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dsn() string {
	username := config.C.GetString("database.sql_username")
	password := config.C.GetString("database.sql_password")
	hostname := config.C.GetString("database.sql_host")
	dbName := config.C.GetString("database.sql_dbname")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	return db, nil
}

func NewRepository(client *sql.DB) src.RepositoryList {
	return &Repository{
		db: client,
	}

}
