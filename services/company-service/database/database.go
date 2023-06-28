package database

import (
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
	"github.com/uptrace/bun"
)

type CompanyDB struct {
	database *bun.DB
}

func NewCompanyDB(database *bun.DB) *CompanyDB {
	return &CompanyDB{
		database: database,
	}
}

type Companies interface {
	CompanyFindOneById(uuid.UUID) (models.Company, int32, error)
	CompanyFindOneByName(string) (models.Company, int32, error)
	CompanyFindAll() ([]models.Company, int32, error)
	CompanyCreateOne(models.Company) (int32, error)
	CompanyUpdateOne(models.Company) (int32, error)
	CompanyDeleteOne(uuid.UUID) (int32, error)
}

type Database struct {
	Company Companies
}
