package postgres

import (
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewConn() *sqlx.DB {
	db, err := sqlx.Open("pgx", os.Getenv("POSTGRES_CONN"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
