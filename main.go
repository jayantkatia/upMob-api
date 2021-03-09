package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jayantkatia/upcoming_mobiles_api/api"
	db "github.com/jayantkatia/upcoming_mobiles_api/db/sqlc"
	"github.com/jayantkatia/upcoming_mobiles_api/scraper"
	"github.com/jayantkatia/upcoming_mobiles_api/util"

	_ "github.com/lib/pq"
)

func prOk() {
	fmt.Print("OK")
}
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	task := func() {
		scraper.StartScraping(store)
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
