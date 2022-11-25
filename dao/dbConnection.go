package dao

import (
	"database/sql"
	"energyByDate/env"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	var (
		host     = env.Env("DB_HOST")
		portStr  = env.Env("DB_PORT")
		user     = env.Env("DB_USER")
		password = env.Env("DB_PASSWORD")
		dbname   = env.Env("DB_NAME")
	)

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
