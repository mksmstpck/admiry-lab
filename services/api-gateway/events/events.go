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

type Users interface {
	UserGetByEmail(email string) (models.User, int32, error)
	UserGetByUsername(username string) (models.User, int32, error)
	UserGetById(id uuid.UUID) (models.User, int32, error)
	UserPost(user *models.User) (models.User, int32, error)
	UserPut(user *models.User) (int32, error)
	UserDelete(id uuid.UUID) (int32, error)
}
