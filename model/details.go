package model

type Detail struct {
	ID          string     `json:"id,omitempty" bson:"_id"`
	ProductName string     `json:"product-name" bson:"product-name"`
	Group       [][]string `json:"group" bson:"group"`
}
