package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSrc = "postgresql://localhost/bank?user=njeri&password=KelynPNjeri@1998"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSrc)
	if err != nil {
		log.Fatal("DB Connection failed: ", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}