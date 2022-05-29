package inventory_repo

import (
	"context"
	"inventoryservice/model"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Repo class
type Repo struct {
	// Connection pool for read-only connection
	dbR *pgxpool.Pool
	// Connection pool for read-write connection
	dbW *pgxpool.Pool
}

// New will return object of Repo class
func New(dbRead, dbWrite *pgxpool.Pool) *Repo {
	return &Repo{
		dbR: dbRead,
		dbW: dbWrite,
	}
}

func (r Repo) GetStock(ctx context.Context, productID int64) (model.Inventory, error) {
	// TODO: Get data from DB using r.dbR connection pool
	return model.Inventory{
		ProductID: productID,
		Stock:     10,
	}, nil
}
