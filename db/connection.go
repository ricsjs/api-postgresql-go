package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	//conf := configs.GetDB()

	sc := "host=localhost port=5432 user=user_todo dbname=api_todo sslmode=disable"

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
