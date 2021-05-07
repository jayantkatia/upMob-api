package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jayantkatia/upcoming_mobiles_api/api"
	db "github.com/jayantkatia/upcoming_mobiles_api/db/sqlc"
	"github.com/jayantkatia/upcoming_mobiles_api/scraper"
	"github.com/jayantkatia/upcoming_mobiles_api/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}

	queries := db.New(conn)
	server := api.NewServer(queries)

	task := func() {
		scraper.Scraper(queries)
	}
	go func() {
		s := gocron.NewScheduler(time.UTC)
		s.Every(24).Hours().Do(task)
		s.StartBlocking()
	}()

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
