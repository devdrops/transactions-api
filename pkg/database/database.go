package database

import (
	"database/sql"
	"fmt"

	"transactions-api/pkg/config"

	_ "github.com/lib/pq"
)

var (
	database string = "postgres"
	connStr  string = "%s://%s:%s@%s/%s?sslmode=%s"
)

type Database struct {
	Conn *sql.DB
}

func New() (*Database, error) {
	c := config.NewConfig()
	cs := fmt.Sprintf(connStr, database, c.DbUser, c.DbPass, c.DbHost, c.DbName, c.DbSSLM)

	db, err := sql.Open(database, cs)
	if err != nil {
		return &Database{}, err
	}

	dbase := Database{
		Conn: db,
	}

	return &dbase, nil
}
