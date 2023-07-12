package handlers

import (
	"github.com/mkskstpck/admiry-lab/pkg/cache"
	"github.com/mkskstpck/admiry-lab/services/perm-service/database"
	"github.com/nats-io/nats.go"
)

type Handler struct {
	conn  *nats.EncodedConn
	perm  database.Perms
	cache cache.Cache
}

func NewHandler(conn *nats.EncodedConn, perm *database.PermDB, cache *cache.Cacher) *Handler {
	return &Handler{
		conn:  conn,
		perm:  perm,
		cache: cache,
	}
}

func (h *Handler) HandleAll() {
	h.permReadById()
	h.permReadByName()
	h.permReadAll()
	h.permCreate()
	h.permUpdate()
	h.permDelete()
}
