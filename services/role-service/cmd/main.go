package main

import (
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/to-rename/pkg/conectors"
	"github.com/mkskstpck/to-rename/services/role-service/config"
	"github.com/mkskstpck/to-rename/services/role-service/database"
	"github.com/mkskstpck/to-rename/services/role-service/handlers"
)

func main() {
	// config
	config := config.NewConfig()

	//nats connection
	c, err := conectors.NewNats(config.NatsURI)
	if err != nil {
		panic(err)
	}

	// postgres connection
	db, err := conectors.NewPsql(
		config.PSQLaddr,
		config.PSQLuser,
		config.PSQLpass,
		config.PSQLdb,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// redis connection
	ccache := conectors.NewCache(
		config.RedisHost,
		config.RedisPort,
		config.RedisDB,
		time.Second*time.Duration(config.RedisExpires),
	)

	// handle requests
	role := database.NewRoleDB(db)
	handlers.NewHandler(c, role, ccache).HandleAll()

	log.Info("company-service is running")

	<-make(chan int)
}
