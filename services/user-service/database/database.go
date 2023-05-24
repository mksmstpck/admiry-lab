package database

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mkskstpck/to-rename/pkg/models"
)

type UserDB struct {
	database *sql.DB
}

func NewUserDB(database *sql.DB) *UserDB {
	return &UserDB{
		database: database,
	}
}

type Users interface {
	UserFindOneById(uuid.UUID) (models.User, int32, error)
	UserFindOneByUsername(string) (models.User, int32, error)
	UserFindOneByEmail(string) (models.User, int32, error)
	UserCreateOne(models.User) (int32, error)
	UserUpdateOne(models.User) (int32, error)
	UserDeleteOne(uuid.UUID) (int32, error)
}

type Database struct {
	User Users
}
