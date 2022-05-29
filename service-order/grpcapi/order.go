package grpcapi

import (
	"context"

	pb "orderservice/proto/order"
)

func (s *Server) GetOrder(ctx context.Context, orderID *pb.OrderID) (*pb.OrderResponse, error) {

	order, err := s.orderSvc.GetOrder(ctx, orderID.GetValue())
	if err != nil {
		return &pb.OrderResponse{}, err
	}

	var products []*pb.Product
	for _, p := range order.Products {
		products = append(products, &pb.Product{
			ProductId: p.ID,
			Name:      p.Name,
		})
	}

	return &pb.OrderResponse{
		Invoice:  order.Invoice,
		Products: products,
		Id:       order.ID,
	}, nil
}
