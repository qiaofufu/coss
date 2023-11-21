package rpc

import (
	"context"
	"coss/pb/meta"
	"coss/pkg/pg"
	"coss/service/meta/internal/controller"
)

func (m *MetaService) Exist(ctx context.Context, req *meta.ExistReq) (*meta.ExistResp, error) {
	var (
		ctrl = controller.NewMetaController(pg.Client)
	)
	result := ctrl.Exist(req.Bucket, req.Key)
	return &meta.ExistResp{Result: result}, nil
}
