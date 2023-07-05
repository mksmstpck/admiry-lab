package main

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/conectors"
	"github.com/mkskstpck/admiry-lab/services/api-gateway/config"
	"github.com/mkskstpck/admiry-lab/services/api-gateway/events"
	"github.com/mkskstpck/admiry-lab/services/api-gateway/handlers"
)

func main() {
	// config
	config := config.NewConfig()

	// nats connection
	conn, err := conectors.NewNats(config.NatsUrl)
	if err != nil {
		log.Fatal(err)
	}

	// starts gin
	router := gin.Default()
	userEvent := events.NewUserEvent(conn)
	companyEvent := events.NewCompanyEvent(conn)
	roleEvent := events.NewRoleEvent(conn)
	handlers.NewHandlers(
		router,
		conn,
		userEvent,
		companyEvent,
		roleEvent).All()
	router.Run(config.EchoUrl)
}
