package events

import (
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/nats-io/nats.go"
	"github.com/pborman/uuid"
)

type User struct {
	conn *nats.EncodedConn
}

func NewUserEvent(conn *nats.EncodedConn) *User {
	return &User{
		conn: conn,
	}
}

type Company struct {
	conn *nats.EncodedConn
}

func NewCompanyEvent(conn *nats.EncodedConn) *Company {
	return &Company{
		conn: conn,
	}
}

type Role struct {
	conn *nats.EncodedConn
}

func NewRoleEvent(conn *nats.EncodedConn) *Role {
	return &Role{
		conn: conn,
	}
}

type Users interface {
	UserGetByEmail(email string) (models.User, int32, error)
	UserGetByUsername(username string) (models.User, int32, error)
	UserGetById(id uuid.UUID) (models.User, int32, error)
	UserPost(user *models.User) (models.User, int32, error)
	UserPut(user *models.User) (int32, error)
	UserDelete(id uuid.UUID) (int32, error)
}

type Companies interface {
	CompanyGetById(id uuid.UUID) (models.Company, int32, error)
	CompanyGetByName(name string) (models.Company, int32, error)
	CompanyGetAll() ([]models.Company, int32, error)
	CompanyPost(company *models.Company) (models.Company, int32, error)
	CompanyPut(company *models.Company) (int32, error)
	CompanyDelete(id uuid.UUID) (int32, error)
}

type Roles interface {
	RoleGetById(id uuid.UUID) (models.Role, int32, error)
	RoleGetByName(name string) (models.Role, int32, error)
	RolePost(role *models.Role) (models.Role, int32, error)
	RolePut(role *models.Role) (int32, error)
	RoleDelete(id uuid.UUID) (int32, error)
}
