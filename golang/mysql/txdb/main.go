package main

import (
	"database/sql"
	"log"

	txdb "github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// we register an sql driver named "txdb"
	txdb.Register("txdb", "mysql", "root:root@tcp(127.0.0.1:3306)/test")
}

func main() {
	// dsn serves as an unique identifier for connection pool
	db, err := sql.Open("txdb", "txdb_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`INSERT INTO t_users(name) VALUES("gopher")`); err != nil {
		log.Fatal(err)
	}
}
