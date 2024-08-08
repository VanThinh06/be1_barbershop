package main

import (
	"barbershop/src/config"
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

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// go:generate statik -src=barbershop/src/shared/doc/swagger -dest=barbershop/src/shared/doc

func main() {
	logger := config.InitLogger()

	logger.Info("Starting application...")

	configSetting, err := config.InitConfig(".")
	if err != nil {
		logger.WithError(err).Fatal("Cannot load config")
	}

	pgxpoolConfig, err := pgxpool.ParseConfig(configSetting.DBSource)
	if err != nil {
		logger.WithError(err).Fatal("Unable to parse database configuration")
	}

	ctx := context.Background()
	conn, err := pgxpool.NewWithConfig(ctx, pgxpoolConfig)
	if err != nil {
		logger.WithError(err).Fatal("Cannot connect to database")
	}
	defer conn.Close()

	firebase, cld, err := config.InitServices(ctx, *configSetting)
	if err != nil {
		logger.WithError(err).Fatal("Failed to initialize Services app")
	}

	store := db.NewStore(conn)
	serviceApp := db.NewServiceApp(firebase, cld)

	go runGatewayServer(*configSetting, store, serviceApp)
	runGrpcServer(configSetting, store, serviceApp)
}

func runGatewayServer(config config.Config, store db.StoreMain, serviceApp db.ServiceApp) {
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
