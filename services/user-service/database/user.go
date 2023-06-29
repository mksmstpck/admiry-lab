package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/mkskstpck/to-rename/pkg/services"
	"github.com/pborman/uuid"
)

func (d *UserDB) UserFindOneById(ID uuid.UUID) (models.User, int32, error) {
	user := models.User{}
	err := d.database.
		NewSelect().
		Model(&user).
		ExcludeColumn("password").
		Where("id = ?", ID).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: user not found by id")
			return models.User{}, 404, errors.New("user not found")
		}
		log.Error("database: ", err)
		return models.User{}, 500, err
	}
	log.Info("database: user found by id")
	return user, 200, nil
}

func (d *UserDB) UserFindOneByEmail(email string) (models.User, int32, error) {
	user := models.User{}
	err := d.database.
		NewSelect().
		Model(&user).
		ExcludeColumn("password").
		Where("email = ?", email).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: user not found by email")
			return models.User{}, 404, errors.New("user not found")
		}
		log.Error("database: ", err)
		return models.User{}, 500, err
	}
	log.Info("database: user found by email")
	return user, 200, nil
}

func (d *UserDB) UserFindOneByUsername(username string) (models.User, int32, error) {
	user := models.User{}
	err := d.database.NewSelect().
		Model(&user).ExcludeColumn("password").
		Where("username = ?", username).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: user not found by username")
			return models.User{}, 404, errors.New("user not found")
		}
		log.Error("database: ", err)
		return models.User{}, 500, err
	}
	log.Info("database: user found by username")
	return user, 200, nil
}

func (d *UserDB) UserCreateOne(user models.User) (int32, error) {
	user.Password = services.PasswordHash(user.Password)
	_, err := d.database.
		NewInsert().
		Model(&user).
		Exec(context.Background())
	if err != nil {
		log.Error("database: ", err)
		return 500, err
	}
	log.Info("database: user created")
	return 201, nil
}

func (d *UserDB) UserUpdateOne(user models.User) (int32, error) {
	res, err := d.database.
		NewUpdate().
		Model(&user).
		ExcludeColumn("id", "password").
		Where("id = ?", user.ID).
		Exec(context.Background())
	if err != nil {
		log.Error("database: ", err)
		return 500, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error("database: ", err)
		return 500, err
	}
	if count == 0 {
		log.Info("database: user not found")
		return 404, errors.New("user not found")
	}
	log.Info("database: user updated")
	return 204, nil
}

func (d *UserDB) UserDeleteOne(ID uuid.UUID) (int32, error) {
	res, err := d.database.
		NewDelete().
		Model(&models.User{ID: ID}).
		Where("id = ?", ID).
		Exec(context.Background())
	if err != nil {
		log.Error("database: ", err)
		return 500, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error("database: ", err)
		return 500, err
	}
	if count == 0 {
		log.Error("database: user not found")
		return 404, errors.New("user not found")
	}
	log.Info("database: user deleted")
	return 204, nil
}
