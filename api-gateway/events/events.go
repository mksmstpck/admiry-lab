package events

import (
	"github.com/mksmstpck/to-rename/api-gateway/models"
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
	UserEmailGet(email string) (models.User, error)
	UserUsernameGet(username string) (models.User, error)
	UserIdGet(id int32) (models.User, error)
	UserPost(user *models.User) (models.User, error)
	UserPut(user *models.User) error
	UserDelete(id int32) error
}
