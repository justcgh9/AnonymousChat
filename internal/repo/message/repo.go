package message

import (
	"github.com/jmoiron/sqlx"
)

const schema = `
CREATE TABLE IF NOT EXISTS message (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
`

type MsgRepo struct {
	DB *sqlx.DB
}

func NewRepo(db *sqlx.DB) *MsgRepo {
	if db == nil {
		panic("null db pointer")
	}

	repo := &MsgRepo{DB: db}
	err := repo.DB.Ping()
	if err != nil {
		panic(err)
	}

	db.MustExec(schema)

	return repo
}
