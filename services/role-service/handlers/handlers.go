package handlers

import (
	"github.com/mkskstpck/admiry-lab/pkg/cache"
	"github.com/mkskstpck/admiry-lab/services/role-service/database"
	"github.com/nats-io/nats.go"
)

type Handler struct {
	conn  *nats.EncodedConn
	role  database.Roles
	cache cache.Cache
}

func NewHandler(conn *nats.EncodedConn, role *database.RoleDB, cache *cache.Cacher) *Handler {
	return &Handler{
		conn:  conn,
		role:  role,
		cache: cache,
	}
}

func (h *Handler) HandleAll() {
	h.roleReadById()
	h.roleReadByName()
	h.roleCreate()
	h.roleUpdate()
	h.roleDelete()
}
