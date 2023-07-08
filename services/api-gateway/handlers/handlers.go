package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"

	"github.com/mkskstpck/admiry-lab/services/api-gateway/events"
)

type Handlers struct {
	g       *gin.Engine
	conn    *nats.EncodedConn
	user    events.Users
	company events.Companies
	role    events.Roles
	perm    events.Perms
}

func NewHandlers(gin *gin.Engine,
	c *nats.EncodedConn,
	user *events.User,
	company *events.Company,
	role *events.Role,
	perm *events.Perm,
) *Handlers {
	return &Handlers{
		g:       gin,
		conn:    c,
		user:    user,
		company: company,
		role:    role,
		perm:    perm,
	}
}

func (h *Handlers) All() {
	// grougs
	user := h.g.Group("/users")
	company := h.g.Group("/companies")
	role := h.g.Group("/roles")
	perm := h.g.Group("/perms")
	// user endpoints
	user.POST("/", h.userCreate)
	user.GET("/id/:id", h.userReadById)
	user.GET("/username/:username", h.userReadByUsername)
	user.GET("/email/:email", h.userReadByEmail)
	user.PUT("/", h.userUpdate)
	user.DELETE("/:id", h.userDelete)
	// company endpoints
	company.POST("/", h.companyCreate)
	company.GET("/id/:id", h.companyReadById)
	company.GET("/name/:name", h.companyReadByName)
	company.GET("/", h.companyReadAll)
	company.PUT("/", h.companyUpdate)
	company.DELETE("/:id", h.companyDelete)
	// role endpoints
	role.POST("/", h.roleCreate)
	role.GET("/id/:id", h.roleReadById)
	role.GET("/name/:name", h.roleReadByName)
	role.PUT("/", h.roleUpdate)
	role.DELETE("/:id", h.roleDelete)
	// permission endpoints
	perm.POST("/", h.permCreate)
	perm.GET("/id/:id", h.permReadById)
	perm.GET("/name/:name", h.permReadByName)
	perm.PUT("/", h.permUpdate)
	perm.DELETE("/:id", h.permDelete)
}
