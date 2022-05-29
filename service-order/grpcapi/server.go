package grpcapi

import (
	"context"
	"log"
	"net"

	"orderservice/model"
	pb "orderservice/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IOrderService interface {
	GetOrder(ctx context.Context, orderID int64) (model.Order, error)
}

type Server struct {
	pb.UnimplementedOrderServer
	server   *grpc.Server
	orderSvc IOrderService
}

// New will create the gRPC Server object.
// Params:
// @o = Order Service implementation
func New(o IOrderService) *Server {
	return &Server{
		orderSvc: o,
	}
}

func (s *Server) Serve(port string) {
	// Create port listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer()
	pb.RegisterOrderServer(s.server, s)
	reflection.Register(s.server)

	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
