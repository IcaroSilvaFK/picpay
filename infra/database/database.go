package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	Provider string
	DSN      string
}

type Queryer interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
	Prepare(string) (*sql.Stmt, error)
	Exec(string, ...interface{}) (sql.Result, error)
}

func NewDbConnection(config Config) *sql.DB {

	db, err := sql.Open(config.Provider, config.DSN)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db

}
