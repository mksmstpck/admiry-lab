package database

import (
	"database/sql"
	"errors"

	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/mkskstpck/to-rename/pkg/services"
)

func (d *UserDB) UserFindOneId(ID int32) (models.User, int32, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email", "roleid" from "Users" where "id" = $1`
	row := d.database.QueryRow(selectQuery, ID)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.RoleID)
	if err != sql.ErrNoRows {
		return user, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserFindOneEmail(email string) (models.User, int32, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email", "roleid" from "Users" where "email" = $1`
	row := d.database.QueryRow(selectQuery, email)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.RoleID)
	if err != sql.ErrNoRows {
		return user, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserFindOneUsername(username string) (models.User, int32, error) {
	var user models.User
	selectQuery := `select "id", "username", "fullname", "email", "roleid" from "Users" where "username" = $1`
	row := d.database.QueryRow(selectQuery, username)
	err := row.Scan(&user.ID, &user.Username, &user.FullName, &user.Email, &user.RoleID)
	if err != sql.ErrNoRows {
		return user, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserCreateOne(user models.User) (int32, error) {
	insertQuery := `insert into "Users"("username", "fullname", "email", "password", "roleid") values ($1, $2, $3, $4, $5)`
	_, err := d.database.Exec(
		insertQuery,
		user.Username,
		user.FullName,
		user.Email,
		services.PasswordHash(user.Password),
		user.RoleID)
	if err != sql.ErrNoRows {
		return 500, err
	}
	return 200, nil
}

func (d *UserDB) UserUpdateOne(user models.User) (int32, error) {
	u, code, err := d.UserFindOneId(user.ID)
	if u.ID == 0 {
		return 404, errors.New("user not found")
	}
	if code != 200 {
		return code, err
	}
	updateQuery := `update "Users" set "fullname" = $1, "password" = $2, "roleid" = $3 where "id" = $4`
	_, err = d.database.Exec(
		updateQuery,
		user.FullName,
		services.PasswordHash(user.Password),
		user.RoleID,
		user.ID)
	if err != sql.ErrNoRows {
		return 500, err
	}
	return 200, nil
}

func (d *UserDB) UserDeleteOne(ID int32) (int32, error) {
	u, code, err := d.UserFindOneId(ID)
	if u.ID == 0 {
		return 404, errors.New("user not found")
	}
	if code != 200 {
		return code, err
	}
	deleteQuery := `delete from "Users" where "id" = $1`
	_, err = d.database.Exec(deleteQuery, ID)
	if err != sql.ErrNoRows {
		return 500, err
	}
	return 200, nil
}
