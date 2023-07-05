package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/conectors"
	"github.com/mkskstpck/to-rename/services/api-gateway/config"
	"github.com/mkskstpck/to-rename/services/api-gateway/events"
	handlers "github.com/mkskstpck/to-rename/services/api-gateway/handlers/web"
)

func main() {
	// config
	config := config.NewConfig()

	// nats connection
	c, err := conectors.NewNats(config.NatsUrl)
	if err != nil {
		log.Fatal(err)
	}

	// starts echo
	r := gin.Default()
	userEvent := events.NewUserEvent(c)
	companyEvent := events.NewCompanyEvent(c)
	roleEvent := events.NewRoleEvent(c)
	handlers.NewHandlers(
		r,
		c,
		userEvent,
		companyEvent,
		roleEvent).All()
	r.Run(config.EchoUrl)
}
