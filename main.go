package main

import (
	"barbershop/api"
	db "barbershop/db/sqlc"
	"barbershop/utils"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	// Load file config
	config, err := utils.LoadConfig(".") // . vì nằm cùng vị trí với thư mục hiện tại app.env
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// connect possgres
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	//
	queries := db.NewStore(conn)
	server, err := api.NewServer(config, queries)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
