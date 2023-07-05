package handlers

import (
	"github.com/mkskstpck/admiry-lab/pkg/cache"
	"github.com/mkskstpck/admiry-lab/services/user-service/database"
	"github.com/nats-io/nats.go"
)

type Handler struct {
	conn  *nats.EncodedConn
	user  database.Users
	cache cache.Cache
}

func NewHandler(conn *nats.EncodedConn, user *database.UserDB, cache *cache.Cacher) *Handler {
	return &Handler{
		conn:  conn,
		user:  user,
		cache: cache,
	}
}

func (h *Handler) HandleAll() {
	h.userReadById()
	h.userReadByUsername()
	h.userReadByEmail()
	h.userCreate()
	h.userUpdate()
	h.userDelete()
}
