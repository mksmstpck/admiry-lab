package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/mkskstpck/to-rename/user-service/cache"
	"github.com/mkskstpck/to-rename/user-service/config"
	"github.com/mkskstpck/to-rename/user-service/database"
	"github.com/mkskstpck/to-rename/user-service/handlers"
	"github.com/nats-io/nats.go"
)

func main() {
	// config
	config, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	//nats connection
	nc, err := nats.Connect(config.NatsUrl)
	if err != nil {
		panic(err)
	}
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	// postgres connection
	port, err := strconv.Atoi(config.PSQLport)
	if err != nil {
		panic(err)
	}
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PSQLhost,
		port,
		config.PSQLuser,
		config.PSQLpass,
		config.PSQLdb,
	)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Print(err)
	}
	log.Println("postgres connected!")
	// redis connection
	ucache := cache.NewRedisCache(
		config.RedisHost,
		config.RedisPort,
		config.RedisDB,
		time.Second*time.Duration(config.RedisExpires))
	log.Println("redis connected!")
	// handle requests
	user := database.NewUserDB(db)
	handlers.NewHandler(c, user, ucache).HandleAll()

	<-make(chan int)
}
