package grpcapi

import (
	"context"
	pb "productservice/proto/product"
)

func (s *Server) GetProduct(ctx context.Context, productID *pb.ProductID) (*pb.Product, error) {
	// Get data from service
	product, err := s.prodSvc.GetProductInfo(ctx, productID.GetValue())
	if err != nil {
		return &pb.Product{}, err
	}

	return &pb.Product{
		ProductId: product.ID,
		Name:      product.Name,
		Stock:     int32(product.Stock),
		Price:     float64(product.Price),
	}, nil
}
