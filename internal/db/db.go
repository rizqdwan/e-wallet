package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}


	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}
