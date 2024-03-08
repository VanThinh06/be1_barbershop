package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://mtt16:Vanthinh11@localhost:5432/BarberShop?sslmode=disable"
)

func TestMain(m *testing.M) {
	_, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// testQueries = New(conn)
	os.Exit(m.Run())
}
