package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (d *RoleDB) RoleFindOneById(id uuid.UUID) (models.Role, int32, error) {
	role := models.Role{}
	err := d.database.NewSelect().Model(&role).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: role not found")
			return models.Role{}, 404, errors.New("role not found")
		}
		log.Error("database: ", err)
		return models.Role{}, 500, err
	}
	log.Info("database: company found")
	return role, 200, nil
}

func (d *RoleDB) RoleFindOneByName(name string) (models.Role, int32, error) {
	role := models.Role{}
	err := d.database.NewSelect().Model(&role).Where("name = ?", name).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: role not found")
			return models.Role{}, 404, errors.New("role not found")
		}
		log.Error("database: ", err)
		return models.Role{}, 500, err
	}
	log.Info("database: company found")
	return role, 200, nil
}

func (d *RoleDB) RoleCreateOne(role models.Role) (int32, error) {
	_, err := d.database.NewInsert().Model(&role).Exec(context.Background())
	if err != nil {
		log.Print(role)
		log.Error("database: ", err)
		return 500, err
	}
	log.Info("database: role created")
	return 200, nil
}

func (d *RoleDB) RoleUpdateOne(role models.Role) (int32, error) {
	res, err := d.database.NewUpdate().Model(&role).Where("id = ?", role.ID).Exec(context.Background())
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
		log.Info("database: role not found")
		return 404, errors.New("role not found")
	}
	log.Info("database: role updated")
	return 204, nil
}

func (d *RoleDB) RoleDeleteOne(id uuid.UUID) (int32, error) {
	res, err := d.database.NewDelete().Model(&models.Role{}).Where("id = ?", id).Exec(context.Background())
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
		log.Info("database: role not found")
		return 404, errors.New("role not found")
	}
	log.Info("database: role deleted")
	return 200, nil
}
