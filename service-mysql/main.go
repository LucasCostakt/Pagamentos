package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"service-mysql/repository"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "admin"
	hostname = "mysql:3306"
	dbname   = "mydb"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	log.Printf("Connected to DB %s successfully\n", dbname)

	var repo = repository.NewRepository(db)

	err = repo.CreateUserTypeTable()
	if err != nil {
		log.Println("Error CreateUserTypeTable() ", err)
		return
	}

	err = repo.CreateUserTable()
	if err != nil {
		log.Println("Error CreateUserTable() ", err)
		return
	}
	err = repo.CreateTransactionTable()
	if err != nil {
		log.Println("Error CreateTransactionTable() ", err)
		return
	}
	err = repo.CreateReversalTable()
	if err != nil {
		log.Println("Error CreateReversalTable() ", err)
		return
	}
	err = repo.CreateUsersTypes()
	if err != nil {
		log.Println("Error CreateUsersTypes() ", err)
		return
	}
	err = repo.CreateUsers()
	if err != nil {
		log.Println("Error CreateUsers() ", err)
		return
	}

}
