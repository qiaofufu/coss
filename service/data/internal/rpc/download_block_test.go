package rpc

import (
	"context"
	"coss/pb/data"
	"coss/pkg/test"
	"testing"
)

func TestDataService_DownloadBlock(t *testing.T) {
	var (
		client = test.GetClient("127.0.0.1:3001", data.NewDataClient)
		ctx    = context.TODO()
	)
	resp, err := client.DownloadBlock(ctx, &data.DownloadBlockReq{
		BucketID:  1,
		Offset:    0,
		Limit:     5,
		TableName: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", resp)
}
