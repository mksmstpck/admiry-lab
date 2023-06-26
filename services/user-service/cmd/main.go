package main

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/mkskstpck/to-rename/pkg/conectors"
	"github.com/mkskstpck/to-rename/services/user-service/config"
	"github.com/mkskstpck/to-rename/services/user-service/database"
	"github.com/mkskstpck/to-rename/services/user-service/handlers"
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
		panic(err)
	}
	defer db.Close()

	// redis connection
	ucache := conectors.NewCache(
		config.RedisHost,
		config.RedisPort,
		config.RedisDB,
		time.Second*time.Duration(config.RedisExpires),
	)

	// handle requests
	user := database.NewUserDB(db)
	handlers.NewHandler(c, user, ucache).HandleAll()

	<-make(chan int)
}
