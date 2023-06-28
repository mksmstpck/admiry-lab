package handlers

import (
	"github.com/mkskstpck/to-rename/pkg/cache"
	"github.com/mkskstpck/to-rename/services/company-service/database"
	"github.com/nats-io/nats.go"
)

type Handler struct {
	conn    *nats.EncodedConn
	company database.Companies
	cache   cache.Cache
}

func NewHandler(conn *nats.EncodedConn, company *database.CompanyDB, cache *cache.Cacher) *Handler {
	return &Handler{
		conn:    conn,
		company: company,
		cache:   cache,
	}
}

func (h *Handler) HandleAll() {
	h.companyReadById()
	h.companyReadByName()
	h.companyReadAll()
	h.companyCreate()
	h.companyUpdate()
	h.companyDelete()
}
