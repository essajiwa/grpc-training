package grpcapi

import (
	"context"
	pb "inventoryservice/proto/inventory"
)

func (s *Server) GetStock(ctx context.Context, productID *pb.ProductID) (*pb.Inventory, error) {
	// Get data from service
	inv, err := s.invSvc.GetStock(ctx, productID.GetValue())

	if err != nil {
		return &pb.Inventory{}, err
	}

	return &pb.Inventory{
		ProductId: inv.ProductID,
		Stock:     inv.Stock,
	}, nil
}
