package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/mkskstpck/admiry-lab/pkg/utils"
	"github.com/mkskstpck/admiry-lab/services/api-gateway/events"
)

type Middleware struct {
	jwt             utils.JWTs
	user            events.Users
	accessPublicKey []byte
}

func NewMiddleware(jwt utils.JWT, users *events.User, accessPublicKey []byte) *Middleware {
	return &Middleware{
		jwt:             jwt,
		user:            users,
		accessPublicKey: accessPublicKey,
	}
}

func (m *Middleware) DeserializeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{Message: "Unauthorized"})
			return
		}
		user_id, err := m.jwt.Validate(authHeader, m.accessPublicKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{Message: "Unauthorized"})
			return
		}

		user, code, err := m.user.UserGetById(user_id)
		if err != nil {
			c.AbortWithStatusJSON(int(code), models.Message{Message: err.Error()})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
