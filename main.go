package main

import (
	db "barbershop/src/db/sqlc"
	gapi "barbershop/src/gapi/barber"
	customergapi "barbershop/src/gapi/customer"
	"barbershop/src/pb/barber"
	"barbershop/src/pb/customer"
	"context"
	"log"
	"net"
	"net/http"

	_ "barbershop/src/shared/doc/statik"
	"barbershop/src/shared/utilities"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"

	firebase "firebase.google.com/go/v4"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// go:generate statik -src=barbershop/src/shared/doc/swagger -dest=barbershop/src/shared/doc

func main() {

	config, err := utilities.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("Khởi tạo ứng dụng Firebase thất bại: %v", err)
	}

	store := db.NewStore(conn)

	firebaseApp := db.NewFirebase(app)
	go runGatewayServer(config, store, firebaseApp)
	runGrpcServer(config, store, firebaseApp)
}

func runGatewayServer(config utilities.Config, store db.StoreMain, firebase db.FirebaseApp) {
	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server, err := gapi.NewServer(config, store, firebase)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = barber.RegisterBarberServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server:", err)
	}

	serverCustomer, err := customergapi.NewServer(config, store, firebase)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = customer.RegisterCustomerBarberShopHandlerServer(ctx, grpcMux, serverCustomer)
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

func runGrpcServer(config utilities.Config, store db.StoreMain, firebase db.FirebaseApp) {
	grpcServer := grpc.NewServer()

	server, err := gapi.NewServer(config, store, firebase)
	if err != nil {
		log.Fatal("cannot create service:", err)
	}
	barber.RegisterBarberServiceServer(grpcServer, server)

	serverCustomer, err := customergapi.NewServer(config, store, firebase)
	if err != nil {
		log.Fatal("cannot create service Customer:", err)
	}
	customer.RegisterCustomerBarberShopServer(grpcServer, serverCustomer)

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
