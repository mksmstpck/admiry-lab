package database

import (
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/pborman/uuid"
	"github.com/uptrace/bun"
)

type RoleDB struct {
	database *bun.DB
}

func NewRoleDB(database *bun.DB) *RoleDB {
	return &RoleDB{
		database: database,
	}
}

type Roles interface {
	RoleFindOneById(uuid.UUID) (models.Role, int32, error)
	RoleFindOneByName(string) (models.Role, int32, error)
	RoleCreateOne(models.Role) (int32, error)
	RoleUpdateOne(models.Role) (int32, error)
	RoleDeleteOne(uuid.UUID) (int32, error)
}

type Database struct {
	Role Roles
}
