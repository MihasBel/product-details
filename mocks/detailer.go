package mocks

import (
	"github.com/MihasBel/product-details/model"
)

type FakeDetailer struct{}

func (FakeDetailer) All() ([]model.Detail, error) {
	var testDetails []model.Detail = []model.Detail{
		model.Detail{
			ID:          "123",
			ProductName: "",
			Group:       nil,
		},
		model.Detail{
			ID:          "321",
			ProductName: "",
			Group:       nil,
		},
		model.Detail{
			ID:          "1010",
			ProductName: "",
			Group:       nil,
		},
	}
	return testDetails, nil
}

func (FakeDetailer) ByID(id string) (model.Detail, error) {
	//TODO implement me
	panic("implement me")
}

func (FakeDetailer) InsertOne(d model.Detail) (model.Detail, error) {
	//TODO implement me
	panic("implement me")
}
