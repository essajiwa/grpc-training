package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"inventoryservice/repository/inventory_repo"
	"inventoryservice/service/inventory"
	"syscall"

	"inventoryservice/grpcapi"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Start dependency creation to run this service
	var (
		grpcPort        string = ":50053"
		grpcServer      *grpcapi.Server
		readDB, writeDB string
	)

	if os.Getenv("GRPC_PORT") != "" {
		grpcPort = os.Getenv("GRPC_PORT")
	}

	// DB Initialization
	readDB = os.Getenv("DB_READ_URL")
	if readDB == "" {
		log.Fatal("Please set DB_READ_URL environment variable")
	}

	writeDB = os.Getenv("DB_WRITE_URL")
	if writeDB == "" {
		log.Fatal("Please set DB_WRITE_URL environment variable")
	}

	// Pool config
	readCfg, err := pgxpool.ParseConfig(readDB)
	if err != nil {
		log.Fatalln("Cant parse configuration for DB_READ_URL")
	}
	readCfg.MaxConns = 50
	readCfg.MaxConnIdleTime = 500

	poolRead, err := pgxpool.ConnectConfig(context.Background(), readCfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer poolRead.Close()

	writeCfg, err := pgxpool.ParseConfig(writeDB)
	if err != nil {
		log.Fatalln("Cant parse configuration for DB_WRITE_URL")
	}

	writeCfg.MaxConns = 50
	writeCfg.MaxConnIdleTime = 500

	poolWrite, err := pgxpool.ConnectConfig(context.Background(), writeCfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer poolWrite.Close()

	invRepo := inventory_repo.New(poolRead, poolWrite)
	invSvc := inventory.New(invRepo)
	grpcServer = grpcapi.New(invSvc)

	runServer(*grpcServer, grpcPort)
}

func runServer(server grpcapi.Server, port string) {
	iddleConnsClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// Graceful shutdown
		server.Shutdown()
		log.Println("gRPC server shutdown gracefully")
		close(iddleConnsClosed)
	}()

	log.Println("gRPC server running on port", port)
	server.Serve(port)
	<-iddleConnsClosed
}
