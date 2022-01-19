package dao

import (
	"database/sql"
	"log"
	"time"
)

func openDB() *sql.DB {
	db, err :=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/library")
	if err!=nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		return nil
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.Ping()
	return db
}
