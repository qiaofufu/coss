package rpc

import (
	"context"
	"coss/pb/data"
	"coss/pkg/test"
	"testing"
)

func TestDataService_SaveBlock(t *testing.T) {
	var (
		client = test.GetClient("127.0.0.1:3001", data.NewDataClient)
		ctx    = context.TODO()
		req    = &data.SaveBlockReq{Data: []byte("test data")}
	)
	resp, err := client.SaveBlock(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", resp)
}
