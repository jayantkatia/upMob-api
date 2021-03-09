package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/jayantkatia/backend_upcoming_mobiles/api"
	db "github.com/jayantkatia/backend_upcoming_mobiles/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/upcoming_mobiles?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}
	queries := db.New(conn)

	server := api.NewServer(queries)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
