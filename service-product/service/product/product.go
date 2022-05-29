package product

import (
	"context"

	"productservice/model"
)

type IProductRepository interface {
	GetProductInfo(ctx context.Context, productID int64) (model.Product, error)
}

type ProductService struct {
	product IProductRepository
}

// New will create the ProductService object.
// Params:
// @p = Product Repository
func New(p IProductRepository) *ProductService {
	return &ProductService{
		product: p,
	}
}

func (p ProductService) GetProductInfo(ctx context.Context, productID int64) (model.Product, error) {
	// Put business logic here if needed
	return p.product.GetProductInfo(ctx, productID)
}
