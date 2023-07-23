package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mkskstpck/admiry-lab/pkg/models"
	"github.com/mkskstpck/admiry-lab/services/api-gateway/events"
	"github.com/mkskstpck/admiry-lab/services/api-gateway/jwt"
)

type Middleware struct {
	user          events.Users
	access_secret []byte
}

func NewMiddleware(user events.Users, access_secret []byte) *Middleware {
	return &Middleware{
		user:          user,
		access_secret: access_secret,
	}
}

func (m *Middleware) DeserializeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		user_id, err := jwt.ValidateJWT(token, m.access_secret)

		user, code, err := m.user.UserGetById(user_id)
		if err != nil {
			c.AbortWithStatusJSON(int(code), models.Message{Message: err.Error()})
			return
		}
		c.Set("User", user)
		c.Next()
	}
}
