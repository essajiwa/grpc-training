package inventory

import (
	"context"

	"inventoryservice/model"
)

type IInventoryRepository interface {
	GetStock(ctx context.Context, productID int64) (model.Inventory, error)
}

type InventoryService struct {
	repo IInventoryRepository
}

// New will create the InventoryService object.
// Params:
// @i = Inventory Repository
func New(i IInventoryRepository) *InventoryService {
	return &InventoryService{
		repo: i,
	}
}

func (i InventoryService) GetStock(ctx context.Context, productID int64) (model.Inventory, error) {
	// Put business logic here
	return i.repo.GetStock(ctx, productID)
}
