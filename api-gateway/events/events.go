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
	UserEmailGet(email string) (models.User, int32, error)
	UserUsernameGet(username string) (models.User, int32, error)
	UserIdGet(id int32) (models.User, int32, error)
	UserPost(user *models.User) (models.User, int32, error)
	UserPut(user *models.User) (int32, error)
	UserDelete(id int32) (int32, error)
}
