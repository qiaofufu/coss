package rpc

import (
	"context"
	"coss/pb/meta"
	"coss/pkg/pg"
	"coss/service/meta/internal/controller"
	"github.com/sirupsen/logrus"
)

func (m *MetaService) CreateMeta(ctx context.Context, req *meta.CreateMetaReq) (*meta.CreateMetaResp, error) {
	var (
		ctrl = controller.NewMetaController(pg.Client)
		err  error
	)

	if req.Bucket == "" {
		return &meta.CreateMetaResp{
			Code: 1,
			Msg:  "bucket parameter cannot be empty",
		}, nil
	}

	if req.Key == "" {
		return &meta.CreateMetaResp{
			Code: 1,
			Msg:  "key parameter cannot be empty",
		}, nil
	}

	if req.ObjectID == 0 {
		return &meta.CreateMetaResp{
			Code: 1,
			Msg:  "objectID parameter cannot be 0",
		}, nil
	}

	err = ctrl.CreateMeta(req.Bucket, req.Key, req.ObjectID)
	if err != nil {
		logrus.Error(err)
		return &meta.CreateMetaResp{
			Code: 2,
			Msg:  "create meta failed",
		}, nil
	}

	return &meta.CreateMetaResp{
		Code: 0,
		Msg:  "create meta success",
	}, nil
}
