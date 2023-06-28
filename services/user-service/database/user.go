package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/mkskstpck/to-rename/pkg/services"
	"github.com/pborman/uuid"
)

func (d *UserDB) UserFindOneById(ID uuid.UUID) (models.User, int32, error) {
	user := models.User{}
	err := d.database.NewSelect().Model(&user).Where("id = ?", ID).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, 404, errors.New("user not found")
		}
		return models.User{}, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserFindOneByEmail(email string) (models.User, int32, error) {
	user := models.User{}
	err := d.database.NewSelect().Model(&user).Where("email = ?", email).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, 404, errors.New("user not found")
		}
		return models.User{}, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserFindOneByUsername(username string) (models.User, int32, error) {
	user := models.User{}
	err := d.database.NewSelect().Model(&user).Where("username = ?", username).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, 404, errors.New("user not found")
		}
		return models.User{}, 500, err
	}
	return user, 200, nil
}

func (d *UserDB) UserCreateOne(user models.User) (int32, error) {
	user.Password = services.PasswordHash(user.Password)
	_, err := d.database.NewInsert().Model(&user).Exec(context.Background())
	if err != nil {
		return 500, err
	}
	return 201, nil
}

func (d *UserDB) UserUpdateOne(user models.User) (int32, error) {
	res, err := d.database.NewUpdate().Model(&user).Where("id = ?", user.ID).Exec(context.Background())
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

func (d *UserDB) UserDeleteOne(ID uuid.UUID) (int32, error) {
	res, err := d.database.NewDelete().Model(&models.User{ID: ID}).Where("id = ?", ID).Exec(context.Background())
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
