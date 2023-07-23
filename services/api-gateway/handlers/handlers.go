package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"

	"github.com/mkskstpck/admiry-lab/services/api-gateway/events"
	"github.com/mkskstpck/admiry-lab/services/api-gateway/middleware"
)

type Handlers struct {
	g              *gin.Engine
	conn           *nats.EncodedConn
	user           events.Users
	company        events.Companies
	role           events.Roles
	perm           events.Perms
	access_secret  []byte
	refresh_secret []byte
	access_exp     time.Duration
	refresh_exp    time.Duration
}

func NewHandlers(gin *gin.Engine,
	conn *nats.EncodedConn,
	user *events.User,
	company *events.Company,
	role *events.Role,
	perm *events.Perm,
	access_secret []byte,
	refresh_secret []byte,
	access_exp time.Duration,
	refresh_exp time.Duration,
) *Handlers {
	return &Handlers{
		g:              gin,
		conn:           conn,
		user:           user,
		company:        company,
		role:           role,
		perm:           perm,
		access_secret:  access_secret,
		refresh_secret: refresh_secret,
		access_exp:     access_exp,
		refresh_exp:    refresh_exp,
	}
}
func (h *Handlers) All() {
	// middleware
	authMiddleware := middleware.NewMiddleware(h.user, h.access_secret).DeserializeUser()
	// grougs
	user := h.g.Group("/users")
	company := h.g.Group("/companies")
	role := h.g.Group("/roles")
	perm := h.g.Group("/perms")
	auth := h.g.Group("/auth")
	// user endpoints
	user.POST("/", h.userCreate)
	user.GET("/me", authMiddleware, h.userReadMe)
	user.GET("/id/:id", authMiddleware, h.userReadById)
	user.GET("/username/:username", authMiddleware, h.userReadByUsername)
	user.GET("/email/:email", authMiddleware, h.userReadByEmail)
	user.PUT("/", authMiddleware, h.userUpdate)
	user.DELETE("/:id", authMiddleware, h.userDelete)
	// company endpoints
	company.POST("/", authMiddleware, h.companyCreate)
	company.GET("/id/:id", authMiddleware, h.companyReadById)
	company.GET("/name/:name", authMiddleware, h.companyReadByName)
	company.GET("/", authMiddleware, h.companyReadAll)
	company.PUT("/", authMiddleware, h.companyUpdate)
	company.DELETE("/:id", authMiddleware, h.companyDelete)
	// role endpoints
	role.POST("/", authMiddleware, h.roleCreate)
	role.GET("/id/:id", authMiddleware, h.roleReadById)
	role.GET("/name/:name", authMiddleware, h.roleReadByName)
	role.PUT("/", authMiddleware, h.roleUpdate)
	role.DELETE("/:id", authMiddleware, h.roleDelete)
	// permission endpoints
	perm.POST("/", authMiddleware, h.permCreate)
	perm.GET("/id/:id", authMiddleware, h.permReadById)
	perm.GET("/name/:name", authMiddleware, h.permReadByName)
	perm.GET("/", authMiddleware, h.permReadAll)
	perm.PUT("/", authMiddleware, h.permUpdate)
	perm.DELETE("/:id", authMiddleware, h.permDelete)
	// auth endpoints
	auth.POST("/login-by-email", h.logInByEmail)
	auth.POST("/login-by-username", h.logInByUsername)
	auth.GET("/refresh", authMiddleware, h.refresh)
}
