package main

import (
	db "barbershop/src/db/sqlc"
	gapi "barbershop/src/gapi/barber"
	customergapi "barbershop/src/gapi/customer"
	"barbershop/src/pb/barber"
	"barbershop/src/pb/customer"
	"barbershop/src/utils"
	"context"
	"log"
	"net"
	"net/http"

	_ "barbershop/src/shared/doc/statik"

	"github.com/cloudinary/cloudinary-go/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	firebase "firebase.google.com/go/v4"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// go:generate statik -src=barbershop/src/shared/doc/swagger -dest=barbershop/src/shared/doc

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	pgxpoolConfig, err := pgxpool.ParseConfig(config.DBSource)
	if err != nil {
		log.Fatalf("Unable to parse database configuration: %v", err)
	}
	ctx := context.Background()
	conn, err := pgxpool.NewWithConfig(ctx, pgxpoolConfig)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	firebase, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("Khởi tạo ứng dụng Firebase thất bại: %v", err)
	}

	cld, _ := cloudinary.NewFromParams(config.CloudinaryName, config.CloudinaryKey, config.CloudinarySecret)

	store := db.NewStore(conn)

	serviceApp := db.NewServiceApp(firebase, cld)

	go runGatewayServer(config, store, serviceApp)
	runGrpcServer(config, store, serviceApp)
}

func runGatewayServer(config utils.Config, store db.StoreMain, serviceApp db.ServiceApp) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := runtime.NewServeMux()
	transportCredentials := insecure.NewCredentials()
	_ = []grpc.DialOption{grpc.WithTransportCredentials(transportCredentials)}

	server, err := gapi.NewServer(config, store, serviceApp)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = barber.RegisterBarberServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server:", err)
	}

	serverCustomer, err := customergapi.NewServer(config, store, serviceApp)
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

func runGrpcServer(config utils.Config, store db.StoreMain, serviceApp db.ServiceApp) {
	grpcServer := grpc.NewServer()

	server, err := gapi.NewServer(config, store, serviceApp)
	if err != nil {
		log.Fatal("cannot create service:", err)
	}
	barber.RegisterBarberServiceServer(grpcServer, server)

	serverCustomer, err := customergapi.NewServer(config, store, serviceApp)
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
