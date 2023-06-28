package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"

	"github.com/mkskstpck/to-rename/services/api-gateway/events"
)

type Handlers struct {
	e       *echo.Echo
	conn    *nats.EncodedConn
	user    events.Users
	company events.Companies
}

func NewHandlers(echo *echo.Echo, c *nats.EncodedConn, user *events.User, company *events.Company) *Handlers {
	return &Handlers{
		e:       echo,
		conn:    c,
		user:    user,
		company: company,
	}
}

func (h *Handlers) All() {
	// grougs
	user := h.e.Group("/users")
	company := h.e.Group("/companies")
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
}
