package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/pborman/uuid"
)

func (d *CompanyDB) CompanyFindOneById(ID uuid.UUID) (models.Company, int32, error) {
	company := models.Company{}
	err := d.database.NewSelect().Model(&company).Where("id = ?", ID).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: company not found")
			return models.Company{}, 404, errors.New("company not found")
		}
		log.Error("database: ", err)
		return models.Company{}, 500, err
	}
	log.Info("database: company found")
	return company, 200, nil
}

func (d *CompanyDB) CompanyFindOneByName(name string) (models.Company, int32, error) {
	company := models.Company{}
	err := d.database.NewSelect().Model(&company).Where("name = ?", name).Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("database: company not found")
			return models.Company{}, 404, errors.New("company not found")
		}
		log.Error("database: ", err)
		return models.Company{}, 500, err
	}
	log.Info("database: company found")
	return company, 200, nil
}

func (d *CompanyDB) CompanyFindAll() ([]models.Company, int32, error) {
	company := []models.Company{}
	err := d.database.NewSelect().Model(&company).Scan(context.Background())
	if err != nil {
		log.Error("database: ", err)
		return []models.Company{}, 500, err
	}
	log.Info("database: company found")
	return company, 200, nil
}

func (d *CompanyDB) CompanyCreateOne(company models.Company) (int32, error) {
	_, err := d.database.NewInsert().Model(&company).Exec(context.Background())
	if err != nil {
		log.Error("database: ", err)
		return 500, err
	}
	log.Info("database: company created")
	return 200, nil
}

func (d *CompanyDB) CompanyUpdateOne(company models.Company) (int32, error) {
	res, err := d.database.NewUpdate().Model(&company).Where("id = ?", company.ID).Exec(context.Background())
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
		log.Info("database: company not found")
		return 404, errors.New("company not found")
	}
	log.Info("database: company updated")
	return 204, nil
}

func (d *CompanyDB) CompanyDeleteOne(ID uuid.UUID) (int32, error) {
	company := &models.Company{ID: ID}
	res, err := d.database.NewDelete().Model(company).WherePK().Exec(context.Background())
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
		log.Info("database: company not found")
		return 404, errors.New("company not found")
	}
	log.Info("database: company deleted")
	return 204, nil
}
