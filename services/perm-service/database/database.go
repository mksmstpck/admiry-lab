package database

import (
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/pborman/uuid"
	"github.com/uptrace/bun"
)

type PermDB struct {
	database *bun.DB
}

func NewPermDB(database *bun.DB) *PermDB {
	return &PermDB{
		database: database,
	}
}

type Perms interface {
	PermFindOneById(uuid.UUID) (models.Role, int32, error)
	PermFindOneByName(string) (models.Role, int32, error)
	PermCreateOne(models.Role) (int32, error)
	PermUpdateOne(models.Role) (int32, error)
	PermDeleteOne(uuid.UUID) (int32, error)
}

type Database struct {
	Perm Perms
}
