package events

import (
	"errors"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (c Company) CompanyGetById(id uuid.UUID) (models.Company, int32, error) {
	var res models.Response[models.Company]
	err := c.conn.Request("companies-id-get", id, &res, time.Second)
	if err != nil {
		log.Error("events.CompanyGetById: ", err)
		return models.Company{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.CompanyGetById: ", res.Error)
		return models.Company{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.CompanyGetById: company found")
	return res.Message, res.Status, nil
}

func (c Company) CompanyGetByName(name string) (models.Company, int32, error) {
	var res models.Response[models.Company]
	err := c.conn.Request("companies-name-get", name, &res, time.Second)
	if err != nil {
		log.Error("events.CompanyGetByName: ", err)
		return models.Company{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.CompanyGetByName: ", res.Error)
		return models.Company{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.CompanyGetByName: company found")
	return res.Message, res.Status, nil
}

func (c Company) CompanyGetAll() ([]models.Company, int32, error) {
	var res models.Response[[]models.Company]
	err := c.conn.Request("companies-get", "", &res, time.Second)
	if err != nil {
		log.Error("events.CompanyGetAll: ", err)
		return []models.Company{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.CompanyGetAll: ", res.Error)
		return []models.Company{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.CompanyGetAll: company found")
	return res.Message, res.Status, nil
}

func (c Company) CompanyPost(company *models.Company) (models.Company, int32, error) {
	var res models.Response[models.Company]
	err := c.conn.Request("companies-post", company, &res, time.Second)
	if err != nil {
		log.Error("events.CompanyPost: ", err)
		return models.Company{}, 500, err
	}
	if res.Error != "" {
		log.Error("events.CompanyPost: ", res.Error)
		return models.Company{}, res.Status, errors.New(res.Error)
	}
	log.Info("events.CompanyPost: company created")
	return res.Message, res.Status, nil
}

func (c Company) CompanyPut(company *models.Company) (int32, error) {
	var res models.Response[string]
	err := c.conn.Request("companies-update", company, &res, time.Second)
	if err != nil {
		log.Error("events.CompanyPut: ", err)
		return 500, err
	}
	if res.Error != "" {
		log.Error("events.CompanyPut: ", res.Error)
		return res.Status, errors.New(res.Error)
	}
	log.Info("events.CompanyPut: company updated")
	return res.Status, nil
}

func (c Company) CompanyDelete(id uuid.UUID) (int32, error) {
	var res models.Response[string]
	err := c.conn.Request("companies-delete", id, &res, time.Second)
	if err != nil {
		log.Error("events.CompanyDelete: ", err)
		return 500, err
	}
	if res.Error != "" {
		log.Error("events.CompanyDelete: ", res.Error)
		return res.Status, errors.New(res.Error)
	}
	log.Info("events.CompanyDelete: company deleted")
	return res.Status, nil
}
