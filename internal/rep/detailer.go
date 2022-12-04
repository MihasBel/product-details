package rep

import (
	"context"

	"github.com/MihasBel/product-details/model"
)

// Detailer interface to work with details DB
type Detailer interface {
	All(ctx context.Context) ([]model.Detail, error)
	ByID(ctx context.Context, id string) (model.Detail, error)
	InsertOne(ctx context.Context, d model.Detail) (model.Detail, error)
}
