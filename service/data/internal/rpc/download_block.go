package rpc

import (
	"context"
	"coss/pb/data"
	"coss/pkg/pg"
	"coss/service/data/internal/controller"
	"github.com/sirupsen/logrus"
)

func (d DataService) DownloadBlock(ctx context.Context, req *data.DownloadBlockReq) (*data.DownloadBlockResp, error) {
	var (
		ctrl = controller.NewDataController(pg.Client)
	)

	dat, err := ctrl.GetBlock(req.BucketID, req.TableName, req.Offset, req.Limit)
	if err != nil {
		logrus.Errorf("get block failed, err: %v", err)
		return &data.DownloadBlockResp{
			Data: nil,
			Msg:  "get block failed",
			Code: 2,
		}, nil
	}
	return &data.DownloadBlockResp{
		Data: dat,
		Msg:  "success",
		Code: 0,
	}, nil
}
