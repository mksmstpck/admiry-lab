package events

import (
	"errors"
	"time"

	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
)

func (c Company) CompanyGetById(id uuid.UUID) (models.Company, int32, error) {
	var res models.Response[models.Company]
	err := c.conn.Request("companies-id-get", id, &res, time.Second)
	if err != nil {
		return models.Company{}, 500, err
	}
	if res.Error != "" {
		return models.Company{}, res.Status, errors.New(res.Error)
	}
	return res.Message, res.Status, nil
}

func (c Company) CompanyGetByName(name string) (models.Company, int32, error) {
	var res models.Response[models.Company]
	err := c.conn.Request("companies-name-get", name, &res, time.Second)
	if err != nil {
		return models.Company{}, 500, err
	}
	if res.Error != "" {
		return models.Company{}, res.Status, errors.New(res.Error)
	}
	return res.Message, res.Status, nil
}

func (c Company) CompanyGetAll() ([]models.Company, int32, error) {
	var res models.Response[[]models.Company]
	err := c.conn.Request("companies-get", "", &res, time.Second)
	if err != nil {
		return []models.Company{}, 500, err
	}
	if res.Error != "" {
		return []models.Company{}, res.Status, errors.New(res.Error)
	}
	return res.Message, res.Status, nil
}

func (c Company) CompanyPost(company *models.Company) (models.Company, int32, error) {
	var res models.Response[models.Company]
	err := c.conn.Request("companies-post", company, &res, time.Second)
	if err != nil {
		return models.Company{}, 500, err
	}
	if res.Error != "" {
		return models.Company{}, res.Status, errors.New(res.Error)
	}
	return res.Message, res.Status, nil
}

func (c Company) CompanyPut(company *models.Company) (int32, error) {
	var res models.Response[string]
	err := c.conn.Request("companies-update", company, &res, time.Second)
	if err != nil {
		return 500, err
	}
	if res.Error != "" {
		return res.Status, errors.New(res.Error)
	}
	return res.Status, nil
}

func (c Company) CompanyDelete(id uuid.UUID) (int32, error) {
	var res models.Response[string]
	err := c.conn.Request("companies-delete", id, &res, time.Second)
	if err != nil {
		return 500, err
	}
	if res.Error != "" {
		return res.Status, errors.New(res.Error)
	}
	return res.Status, nil
}
