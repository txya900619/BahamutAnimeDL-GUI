package database

import (
	"database/sql"
	"fmt"
	"github.com/markbates/pkger"
	migrate "github.com/rubenv/sql-migrate"
	"net/http"
	"os"
)

func migrations(db *sql.DB) {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	migrationSource := &migrate.HttpFileSystemMigrationSource{
		FileSystem: http.FileSystem(pkger.Dir(workingDir + "/migrations")),
	}
	num, err := migrate.Exec(db, "sqlite3", migrationSource, migrate.Up)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}
