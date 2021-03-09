package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jayantkatia/upcoming_mobiles_api/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBTestSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
