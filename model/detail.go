package model

// Detail DTO model to represent one product details
type Detail struct {
	ID          string     `json:"id,omitempty" bson:"_id"`
	ProductName string     `json:"product-name" bson:"product-name"`
	Group       [][]string `json:"group" bson:"group"`
}
