package rpc

import (
	"context"
	"coss/pb/meta"
	"coss/pkg/pg"
	"coss/service/meta/internal/controller"
)

func (m *MetaService) GetMeta(ctx context.Context, req *meta.GetMetaReq) (*meta.GetMetaResp, error) {
	var (
		ctrl = controller.NewMetaController(pg.Client)
	)
	if req.Bucket == "" {
		return &meta.GetMetaResp{
			Code: 1,
			Msg:  "bucket parameter cannot be empty",
		}, nil
	}

	if req.Key == "" {
		return &meta.GetMetaResp{
			Code: 1,
			Msg:  "key parameter cannot be empty",
		}, nil
	}
	res, err := ctrl.GetMeta(req.Bucket, req.Key)
	if err != nil {
		return &meta.GetMetaResp{
			Code: 2,
			Msg:  "get meta failed",
		}, nil
	}

	return &meta.GetMetaResp{
		ObjectID: res.ObjectID,
		Code:     0,
		Msg:      "success",
	}, nil
}
