package sqlite

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func NewConn() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", os.Getenv("SQLITE_PATH"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
