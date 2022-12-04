package mocks

import (
	"context"
	"github.com/MihasBel/product-details/model"
)

// FakeDetailer implement Detailer interface for testing
type FakeDetailer struct{}

// All fake for test get all documents from DB
func (FakeDetailer) All(_ context.Context) ([]model.Detail, error) {
	testDetails := []model.Detail{
		{
			ID:          "123",
			ProductName: "",
			Group:       nil,
		},
		{
			ID:          "321",
			ProductName: "",
			Group:       nil,
		},
		{
			ID:          "1010",
			ProductName: "",
			Group:       nil,
		},
	}
	return testDetails, nil
}

// ByID fake get one detail by id from DB
func (FakeDetailer) ByID(_ context.Context, _ string) (model.Detail, error) {
	//TODO implement me
	panic("implement me")
}

// InsertOne fake insert one new detail to DB. Should generate new object id
func (FakeDetailer) InsertOne(_ context.Context, _ model.Detail) (model.Detail, error) {
	//TODO implement me
	panic("implement me")
}
