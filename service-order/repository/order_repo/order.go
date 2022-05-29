package order_repo

import (
	"context"

	"orderservice/model"
	pb "orderservice/proto/product"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
)

// IProductInfo will holds all method / service required by order repository
type IProductInfo interface {
	GetProduct(ctx context.Context, req *pb.ProductID, opts ...grpc.CallOption) (*pb.Product, error)
}

// Repo class
type Repo struct {
	// Connection pool for read-only connection
	dbR *pgxpool.Pool
	// Connection pool for read-write connection
	dbW *pgxpool.Pool
	// gRPC service Client needed by the service
	prodInfo IProductInfo
}

// New will return object of Repo class
func New(pInfo IProductInfo) *Repo {
	return &Repo{
		prodInfo: pInfo,
	}
}

// GetOrder will return all order with product information as well
func (r *Repo) GetOrder(ctx context.Context, orderID int64) (model.Order, error) {

	var (
		order    model.Order
		products []model.Product
	)

	order.ID = orderID
	order.Invoice = "INV/2022/05/19/0001"

	// Let say we get a bunch of productID from order repository,
	// and then we need to get the product information data from gRPC service Product Info.
	// So this line should be a looping to generate array of productID and sent it to the service as param
	prodInfo, err := r.prodInfo.GetProduct(ctx, &pb.ProductID{Value: 1})
	if err != nil {
		return model.Order{}, err
	}

	products = append(products, model.Product{
		Name:  prodInfo.Name,
		Price: float32(prodInfo.Price),
	})

	order.Products = products

	return order, nil
}
