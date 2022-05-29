package grpcapi

import (
	"context"
	"log"
	"net"

	"productservice/model"
	pb "productservice/proto/product"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IProductService interface {
	GetProductInfo(ctx context.Context, productID int64) (model.Product, error)
}

type Server struct {
	pb.UnimplementedProductInfoServer
	server  *grpc.Server
	prodSvc IProductService
}

// New will create the gRPC Server object.
// Params:
// @o = Order Service implementation
func New(p IProductService) *Server {
	return &Server{
		prodSvc: p,
	}
}

func (s *Server) Serve(port string) {
	// Create port listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()
	pb.RegisterProductInfoServer(s.server, s)
	reflection.Register(s.server)

	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
