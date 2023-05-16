package main

import (
	"barbershop/api"
	db "barbershop/db/sqlc"
	"barbershop/db/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".") // . vì nằm cùng vị trí với thư mục hiện tại app.env
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	queries := db.New(conn)
	server, err := api.NewServer(config, queries)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
