package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func ConnectSqlite() *sql.DB {
	if _, err := os.Stat("saveData.sqlite"); err != nil {
		db, err := sql.Open("sqlite3", "./saveData.sqlite")
		if err != nil {
			fmt.Println(err)
		}
		migrations(db)
		return db
	}
	db, err := sql.Open("sqlite3", "./saveData.sqlite")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
