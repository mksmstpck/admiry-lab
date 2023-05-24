package conectors

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq" // TODO:
)

func NewPsql(
	port string,
	host string,
	user string,
	password string,
	database string,
) (*sql.DB, error) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		portInt,
		user,
		password,
		database,
	)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
