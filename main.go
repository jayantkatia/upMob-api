package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/jayantkatia/backend_upcoming_mobiles/api"
	db "github.com/jayantkatia/backend_upcoming_mobiles/db/sqlc"
	"github.com/jayantkatia/backend_upcoming_mobiles/util"
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
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
