package database

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// Database ...
type Database struct {
	*sqlx.DB
}

// Connect ...
func Connect(connectionString string) (db *Database, err error) {
	conn, err := sqlx.Connect("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	return &Database{DB: conn}, nil
}
