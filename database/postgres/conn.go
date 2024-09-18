package postgres

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/stdlib"
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
