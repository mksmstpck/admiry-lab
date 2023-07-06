package main

import (
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mkskstpck/admiry-lab/pkg/conectors"
	"github.com/mkskstpck/admiry-lab/services/perm-service/config"
	"github.com/mkskstpck/admiry-lab/services/perm-service/database"
	"github.com/mkskstpck/admiry-lab/services/perm-service/handlers"
)

func main() {
	// config
	config := config.NewConfig()

	//nats connection
	c, err := conectors.NewNats(config.NatsURI)
	if err != nil {
		log.Fatal(err)
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
	pcache := conectors.NewCache(
		config.RedisHost,
		config.RedisPort,
		config.RedisDB,
		time.Second*time.Duration(config.RedisExpires),
	)

	// handle requests
	role := database.NewPermDB(db)
	handlers.NewHandler(c, role, pcache).HandleAll()

	log.Info("permission-service is running")

	<-make(chan int)
}
