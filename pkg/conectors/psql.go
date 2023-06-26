package conectors

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewPsql(
	addr string,
	user string,
	password string,
	database string,
) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(addr),
		pgdriver.WithUser(user),
		pgdriver.WithPassword(password),
		pgdriver.WithDatabase(database),
		pgdriver.WithInsecure(true),
		pgdriver.WithDatabase(database),
	))
	db := bun.NewDB(sqldb, pgdialect.New())
	return db, nil
}
