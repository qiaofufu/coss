package rpc

import (
	"coss/pb/meta"
	"google.golang.org/grpc"
	"net"
)

type MetaService struct {
	meta.UnimplementedMetaServer
}

func MustStartGrpc(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	meta.RegisterMetaServer(srv, &MetaService{})
	srv.Serve(lis)
}

//func (m *MetaService) mustEmbedUnimplementedMetaServer() {
//	//TODO implement me
//	panic("implement me")
//}
