package db

import (
	"database/sql"
	"embed"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

//go:embed migrations/*.sql
var embedMigrations embed.FS

func DbConnect() {
	sqliteDsn := os.Getenv("SQLITE_DSN")
	var err error
	DB, err = sql.Open("sqlite3", sqliteDsn)
	if err != nil {
		slog.Error("unable to open database", "error", err)
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(err)
	}

	if err := goose.Up(DB, "migrations"); err != nil {
		panic(err)
	}
}
