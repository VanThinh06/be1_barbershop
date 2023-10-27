package main

import (
	"barbershop/api"
	db "barbershop/db/sqlc"
	"barbershop/gapi"
	"barbershop/pb"
	"barbershop/utils"
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	_ "barbershop/docs/statik"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:generate statik -src=./docs/swagger -dest=./docs

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
	store := db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config, store)
}

func runGatewayServer(config utils.Config, store db.StoreMain) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background()) // create context
	defer cancel()                                          // trì hoãn lệnh cancel trước khi exit khỏi func này

	err = pb.RegisterBarberShopHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start HTTP gateway server: ", err)
	}

}

func runGrpcServer(config utils.Config, store db.StoreMain) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBarberShopServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("start gRPC server at %s", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC: ", err)
	}
}

func runGinServer(config utils.Config, store db.StoreMain) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
