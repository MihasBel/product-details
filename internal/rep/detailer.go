package rep

import "github.com/MihasBel/product-details/model"

type Detailer interface {
	All() ([]model.Detail, error)
	ByID(id string) (model.Detail, error)
	InsertOne(d model.Detail) (model.Detail, error)
}
