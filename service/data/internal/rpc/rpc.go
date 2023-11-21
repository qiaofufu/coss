package rpc

import (
	"coss/pb/data"
	"google.golang.org/grpc"
	"net"
)

type DataService struct {
	data.UnimplementedDataServer
}

func MustStartGrpc(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	data.RegisterDataServer(srv, &DataService{})
	srv.Serve(lis)
}
