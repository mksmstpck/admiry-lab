package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/pborman/uuid"
)

func (d *PermDB) PermFindOneById(id uuid.UUID) (models.Permission, int32, error) {
	perm := models.Permission{}
	err := d.
		database.
		NewSelect().
		Model(&perm).
		Where("id = ?", id).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: permission not found")
			return models.Permission{}, 404, errors.New("permission not found")
		}
		log.Error("database: ", err)
		return models.Permission{}, 500, err
	}
	log.Info("database: permission found")
	return perm, 200, nil
}

func (d *PermDB) PermFindOneByName(name string) (models.Permission, int32, error) {
	perm := models.Permission{}
	err := d.
		database.
		NewSelect().
		Model(&perm).
		Where("name = ?", name).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: permission not found")
			return models.Permission{}, 404, errors.New("permission not found")
		}
		log.Error("database: ", err)
		return models.Permission{}, 500, err
	}
	log.Info("database: permission found")
	return perm, 200, nil
}

func (d *PermDB) PermCreateOne(perm models.Permission) (int32, error) {
	_, err := d.
		database.
		NewInsert().
		Model(&perm).
		Exec(context.Background())
	if err != nil {
		log.Print(perm)
		log.Error("database: ", err)
		return 500, err
	}
	log.Info("database: permission created")
	return 200, nil
}

func (d *PermDB) PermUpdateOne(perm models.Permission) (int32, error) {
	res, err := d.
		database.
		NewUpdate().
		Model(&perm).
		Where("id = ?", perm.ID).
		Exec(context.Background())
	if err != nil {
		log.Error("database: ", err)
		return 500, err
	}
	count, err := res.RowsAffected()
	return int32(count), err
}

func (d *PermDB) PermDeleteOne(id uuid.UUID) (int32, error) {
	res, err := d.
		database.
		NewDelete().
		Model(&models.Permission{}).
		Where("id = ?", id).
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
		log.Info("database: permission not found")
		return 404, errors.New("permission not found")
	}
	log.Info("database: permission deleted")
	return 200, nil
}
