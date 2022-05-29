package main

import (
	"log"
	"os"
	"os/signal"
	pb "productservice/proto/inventory"
	"productservice/repository/product_repo"
	"productservice/service/product"
	"syscall"

	"productservice/grpcapi"

	"google.golang.org/grpc"
)

func main() {
	// Start dependency creation to run this service
	var (
		grpcPort            string = ":50052"
		grpcServer          *grpcapi.Server
		inventoryServiceURL string = "localhost:50053"
	)

	if os.Getenv("GRPC_PORT") != "" {
		grpcPort = os.Getenv("GRPC_PORT")
	}

	if os.Getenv("INVENTORY_SERVICE_URL") != "" {
		inventoryServiceURL = os.Getenv("INVENTORY_SERVICE_URL")
	}

	//Client for Product gRPC
	conn, err := grpc.Dial(inventoryServiceURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	inventoryClient := pb.NewInventoryInfoClient(conn)

	prodRepo := product_repo.New(inventoryClient)
	prodSvc := product.New(prodRepo)
	grpcServer = grpcapi.New(prodSvc)

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
