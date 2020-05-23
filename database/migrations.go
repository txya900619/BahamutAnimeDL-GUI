package database

import (
	"database/sql"
	"fmt"
	"github.com/markbates/pkger"
	migrate "github.com/rubenv/sql-migrate"
	"net/http"
)

func migrations(db *sql.DB) {

	migrationSource := &migrate.HttpFileSystemMigrationSource{
		FileSystem: http.FileSystem(pkger.Dir("/migrations")),
	}
	num, err := migrate.Exec(db, "sqlite3", migrationSource, migrate.Up)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
