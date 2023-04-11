package database

import (
	"database/sql"

	"github.com/mkskstpck/to-rename/user-service/models"
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
	UserFindOneId(ID int32) (models.User, int32, error)
	UserFindOneUsername(username string) (models.User, int32, error)
	UserFindOneEmail(email string) (models.User, int32, error)
	UserCreateOne(user models.User) (int32, error)
	UserUpdateOne(user models.User) (int32, error)
	UserDeleteOne(ID int32) (int32, error)
}

type Database struct {
	User Users
}
