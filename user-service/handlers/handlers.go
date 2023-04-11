package handlers

import (
	"github.com/mkskstpck/to-rename/user-service/cache"
	"github.com/mkskstpck/to-rename/user-service/database"
	"github.com/nats-io/nats.go"
)

type Handler struct {
	conn  *nats.EncodedConn
	user  database.Users
	cache cache.Cache
}

func NewHandler(conn *nats.EncodedConn, user *database.UserDB, cache *cache.UserCache) *Handler {
	return &Handler{
		conn:  conn,
		user:  user,
		cache: cache,
	}
}

func (h *Handler) HandleAll() {
	h.userIdRead()
	h.userUsernameRead()
	h.userEmailRead()
	h.userCreate()
	h.userUpdate()
	h.userDelete()
}
