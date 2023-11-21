package rpc

import (
	"context"
	"coss/pb/meta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestMetaService_CreateMeta(t *testing.T) {
	var (
		host = "127.0.0.1:3000"
		ctx  = context.TODO()
	)
	conn, err := grpc.Dial(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatal(err)
	}
	client := meta.NewMetaClient(conn)
	resp, err := client.CreateMeta(ctx, &meta.CreateMetaReq{
		Bucket:   "test_bucket",
		Key:      "test_key",
		ObjectID: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", resp)
}
