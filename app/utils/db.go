package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Client *sql.DB
}

func Get(connStr string) (*DB, error) {
	db, err := get(connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		Client: db,
	}, nil
}

func (d *DB) Close() error {
	return d.Client.Close()
}

func get(connStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
