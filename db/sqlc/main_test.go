package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "postgres"
	dbSrc = "postgresql://localhost/bank?user=njeri&password=KelynPNjeri@1998"
)
var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSrc)
	if err != nil {
		log.Fatal("DB Connection failed: ", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}