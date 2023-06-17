package database

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/mkskstpck/to-rename/pkg/services"
)

func (d *UserDB) UserFindOneById(ID uuid.UUID) (models.User, int32, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email" from "Users" where "id" = $1`
	row := d.database.QueryRow(selectQuery, ID)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email)
	if err != sql.ErrNoRows {
		return user, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserFindOneByEmail(email string) (models.User, int32, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email" from "Users" where "email" = $1`
	row := d.database.QueryRow(selectQuery, email)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email)
	if err != sql.ErrNoRows {
		return user, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserFindOneByUsername(username string) (models.User, int32, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email" from "Users" where "username" = $1`
	row := d.database.QueryRow(selectQuery, username)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email)
	if err != sql.ErrNoRows {
		return user, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserCreateOne(user models.User) (int32, error) {
	insertQuery := `insert into "Users"("username", "fullname", "email", "password") values ($1, $2, $3, $4)`
	_, err := d.database.Exec(
		insertQuery,
		user.Username,
		user.FullName,
		user.Email,
		services.PasswordHash(user.Password),
	)
	if err != sql.ErrNoRows {
		return 500, err
	}
	return 200, nil
}

func (d *UserDB) UserUpdateOne(user models.User) (int32, error) {
	u, code, err := d.UserFindOneById(user.ID)
	if u.ID == uuid.Nil {
		return 404, errors.New("user not found")
	}
	if code != 200 {
		return code, err
	}
	updateQuery := `update "Users" set "fullname" = $1, "password" = $2 where "id" = $3`
	_, err = d.database.Exec(
		updateQuery,
		user.FullName,
		services.PasswordHash(user.Password),
		user.ID)
	if err != nil {
		return 500, err
	}
	return 200, nil
}

func (d *UserDB) UserDeleteOne(ID uuid.UUID) (int32, error) {
	deleteQuery := `delete from "Users" where "id" = $1`
	res, err := d.database.Exec(deleteQuery, ID)
	if err != nil {
		return 500, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return 500, err
	}
	if count == 0 {
		return 404, errors.New("user not found")
	}
	return 204, nil
}
