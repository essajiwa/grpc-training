package product_repo

import (
	"context"

	"productservice/model"
	pb "productservice/proto/inventory"

	"google.golang.org/grpc"
)

// IProductInfo will holds all method / service required by order repository
type IInventory interface {
	GetStock(ctx context.Context, req *pb.ProductID, opts ...grpc.CallOption) (*pb.Inventory, error)
}

// Repo class
type Repo struct {
	// gRPC service Client needed by the service
	prodInventory IInventory
}

// New will return object of Repo class
func New(i IInventory) *Repo {
	return &Repo{
		prodInventory: i,
	}
}

// GetProductInfo will return product information with stock data from inventory service
func (r *Repo) GetProductInfo(ctx context.Context, productID int64) (model.Product, error) {
	inv, err := r.prodInventory.GetStock(ctx, &pb.ProductID{Value: productID})
	if err != nil {
		return model.Product{}, err
	}

	return model.Product{
		ID:    inv.GetProductId(),
		Name:  "Microlax",
		Stock: int(inv.Stock),
	}, nil
}
