package db

import (
	"log"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
)

func TestCreateService(t *testing.T) {
	var endTime = pgtype.Time{
		Microseconds: 44400000,
		Valid:        true,
	}
	log.Print("", endTime)
}
