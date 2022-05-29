package main

import (
	"log"
	pb "orderservice/proto/product"
	"os"
	"os/signal"
	"syscall"

	"orderservice/grpcapi"
	"orderservice/repository/order_repo"
	"orderservice/service/order"

	"google.golang.org/grpc"
)

func main() {
	// Start dependency creation to run this service
	var (
		grpcPort          string = ":50051"
		grpcServer        *grpcapi.Server
		productServiceURL string = "localhost:50052"
	)

	if os.Getenv("GRPC_PORT") != "" {
		grpcPort = os.Getenv("GRPC_PORT")
	}

	if os.Getenv("PRODUCT_SERVICE_URL") != "" {
		productServiceURL = os.Getenv("PRODUCT_SERVICE_URL")
	}

	//Client for Product gRPC
	conn, err := grpc.Dial(productServiceURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	productInfoClient := pb.NewProductInfoClient(conn)

	orderRepo := order_repo.New(productInfoClient)
	orderSvc := order.New(orderRepo)
	grpcServer = grpcapi.New(orderSvc)

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
