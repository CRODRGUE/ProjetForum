package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func DbCon() (db *sql.DB) {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "forum",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		println(err)
		log.Fatal(err)
	}
	return db
}
