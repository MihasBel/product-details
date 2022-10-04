package mocks

import (
	"github.com/MihasBel/product-details/model"
)

// FakeDetailer implement Detailer interface for testing
type FakeDetailer struct{}

// All fake for test get all documents from DB
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

// ByID fake get one detail by id from DB
func (FakeDetailer) ByID(id string) (model.Detail, error) {
	//TODO implement me
	panic("implement me")
}

// InsertOne fake insert one new detail to DB. Should generate new object id
func (FakeDetailer) InsertOne(d model.Detail) (model.Detail, error) {
	//TODO implement me
	panic("implement me")
}
