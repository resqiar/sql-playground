package db

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
	Raw  *sql.DB
)

func connect() {
	DSN := "admin:admin@tcp(172.17.0.2)/world"

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	Raw = db
}

func InitDB() {
	once.Do(func() {
		connect()
	})
}
