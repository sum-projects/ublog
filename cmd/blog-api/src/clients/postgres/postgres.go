package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	Client, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
}
