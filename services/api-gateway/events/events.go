package events

import (
	"github.com/google/uuid"
	"github.com/mkskstpck/to-rename/pkg/models"
	"github.com/nats-io/nats.go"
)

type User struct {
	conn *nats.EncodedConn
}

func NewUserEvent(conn *nats.EncodedConn) *User {
	return &User{
		conn: conn,
	}
}

type Permission struct {
	conn *nats.EncodedConn
}

func NewPermissionEvent(conn *nats.EncodedConn) *Permission {
	return &Permission{
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
	UserEmailGet(email string) (models.User, int32, error)
	UserUsernameGet(username string) (models.User, int32, error)
	UserIdGet(id uuid.UUID) (models.User, int32, error)
	UserPost(user *models.User) (models.User, int32, error)
	UserPut(user *models.User) (int32, error)
	UserDelete(id uuid.UUID) (int32, error)
}
