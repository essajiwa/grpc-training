package grpcapi

import (
	"context"
	"log"
	"net"

	"inventoryservice/model"
	pb "inventoryservice/proto/inventory"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type IInventoryService interface {
	GetStock(ctx context.Context, productID int64) (model.Inventory, error)
}

type Server struct {
	pb.UnimplementedInventoryInfoServer
	server    *grpc.Server
	invSvc    IInventoryService
	zapLogger *zap.Logger
}

// New will create the gRPC Server object.
// Params:
// @o = Inventory Service implementation
func New(i IInventoryService) *Server {
	// Should be injected so we can control the zap config
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	return &Server{
		invSvc:    i,
		zapLogger: logger,
	}
}

func (s *Server) Serve(port string) {
	// Create port listener
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.server = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(s.zapLogger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	pb.RegisterInventoryInfoServer(s.server, s)
	reflection.Register(s.server)

	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
