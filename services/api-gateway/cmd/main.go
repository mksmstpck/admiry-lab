package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mkskstpck/to-rename/pkg/conectors"
	"github.com/mkskstpck/to-rename/services/api-gateway/config"
	"github.com/mkskstpck/to-rename/services/api-gateway/events"
	validator "github.com/mkskstpck/to-rename/services/api-gateway/handlers/validators"
	handlers "github.com/mkskstpck/to-rename/services/api-gateway/handlers/web"
)

func main() {
	// config
	config := config.NewConfig()

	// nats connection
	c, err := conectors.NewNats(config.NatsUrl)
	if err != nil {
		panic(err)
	}

	// starts echo
	e := echo.New()
	e.Validator = validator.NewCustomeValidator()
	userEvent := events.NewUserEvent(c)
	roleEvent := events.NewRoleEvent(c)
	permissionEvent := events.NewPermissionEvent(c)
	handlers.NewHandlers(
		e,
		c,
		userEvent,
		roleEvent,
		permissionEvent).All()
	e.Logger.Fatal(e.Start(config.EchoUrl))
}
