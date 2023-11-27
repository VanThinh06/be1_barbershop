package main

import (
	db "barbershop/src/db/sqlc"
	gapi "barbershop/src/gapi/barber"
	customergapi "barbershop/src/gapi/customer"
	"barbershop/src/pb/barber"
	"barbershop/src/pb/customer"
	"barbershop/src/shared/utils"
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	_ "barbershop/src/shared/doc/statik"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	firebase "firebase.google.com/go/v4"
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

	// Khởi tạo ứng dụng Firebase sử dụng ADC
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("Khởi tạo ứng dụng Firebase thất bại: %v", err)
	}

	//
	store := db.NewStore(conn)
	firebaseApp := db.NewFirebase(app)
	go runGatewayServer(config, store, firebaseApp)
	// go runGrpcServerCustomer(config, store, firebaseApp)
	runGrpcServer(config, store, firebaseApp)
}

func runGatewayServer(config utils.Config, store db.StoreMain, firebase db.FirebaseApp) {
	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background()) // create context
	defer cancel()                                          // trì hoãn lệnh cancel trước khi exit khỏi func này

	server, err := gapi.NewServer(config, store, firebase)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = barber.RegisterBarberShopHandlerServer(ctx, grpcMux, server)
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

func runGrpcServer(config utils.Config, store db.StoreMain, firebase db.FirebaseApp) {
	grpcServer := grpc.NewServer()

	server, err := gapi.NewServer(config, store, firebase)
	if err != nil {
		log.Fatal("cannot create service:", err)
	}
	barber.RegisterBarberShopServer(grpcServer, server)

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

// func runGatewayServerCustomer(config utils.Config, store db.StoreMain, firebase db.FirebaseApp) {
// 	server, err := customergapi.NewServer(config, store, firebase)
// 	if err != nil {
// 		log.Fatal("cannot create server:", err)
// 	}
// 	grpcMux := runtime.NewServeMux()
// 	ctx, cancel := context.WithCancel(context.Background()) // create context
// 	defer cancel()                                          // trì hoãn lệnh cancel trước khi exit khỏi func này

// 	err = pb.RegisterAuthCustomerBarberShopHandlerServer(ctx, grpcMux, server)
// 	if err != nil {
// 		log.Fatal("cannot register handler server:", err)
// 	}

// 	mux := http.NewServeMux()
// 	mux.Handle("/", grpcMux)

// 	statikFS, err := fs.New()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
// 	mux.Handle("/swagger/", swaggerHandler)

// 	listener, err := net.Listen("tcp", config.HTTPServerAddressCustomer)
// 	if err != nil {
// 		log.Fatal("cannot create listener: ", err)
// 	}

// 	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
// 	err = http.Serve(listener, mux)
// 	if err != nil {
// 		log.Fatal("cannot start HTTP gateway server Customer: ", err)
// 	}
// }

// func runGrpcServerCustomer(config utils.Config, store db.StoreMain, firebase db.FirebaseApp) {
// 	server, err := customergapi.NewServer(config, store, firebase)
// 	if err != nil {
// 		log.Fatal("cannot create server:", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	pb.RegisterAuthCustomerBarberShopServer(grpcServer, server)
// 	reflection.Register(grpcServer)

// 	listener, err := net.Listen("tcp", config.GRPCServerAddressCustomer)
// 	if err != nil {
// 		log.Fatal("cannot create listener: ", err)
// 	}

// 	log.Printf("start gRPC server at %s", config.GRPCServerAddressCustomer)
// 	err = grpcServer.Serve(listener)
// 	if err != nil {
// 		log.Fatal("cannot start gRPC: ", err)
// 	}
// }
