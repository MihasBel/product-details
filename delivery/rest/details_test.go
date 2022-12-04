package rest

import (
	"testing"

	"github.com/MihasBel/product-details/internal/app"
	"github.com/MihasBel/product-details/mocks"
	"github.com/MihasBel/product-details/model"
)

var conf app.Configuration = app.Configuration{
	ConnectionString: "",
	Database:         "",
	Collection:       "",
	APIKey:           "",
}
var det = mocks.FakeDetailer{}

var testREST *REST = New(conf, det)

var normalDetails = []model.Detail{
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

func TestAll_NormalValue(t *testing.T) {
	want := normalDetails
	got, err := testREST.d.All(nil)
	if err != nil {
		t.Errorf("got error while execute ALL()")
		return
	}
	if len(want) != len(got) {
		t.Errorf("result length %v expected %v", len(got), len(want))
		return
	}
	for i, detail := range got {
		if detail.ID != want[i].ID {
			t.Errorf("mismatch found want: %v got: %v", want[i], detail)
			return
		}
	}
}
