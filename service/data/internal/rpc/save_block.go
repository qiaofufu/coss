package rpc

import (
	"context"
	"coss/pb/data"
	"coss/pkg/pg"
	"coss/service/data/internal/controller"
	"github.com/sirupsen/logrus"
)

func (d DataService) SaveBlock(ctx context.Context, req *data.SaveBlockReq) (*data.SaveBlockResp, error) {
	var (
		ctrl = controller.NewDataController(pg.Client)
	)
	block, err := ctrl.SaveBlock(req.GetData())
	if err != nil {
		logrus.Errorf("save block failed, err: %v", err)
		return &data.SaveBlockResp{
			Block: nil,
			Msg:   "save block failed",
			Code:  2,
		}, nil
	}
	return &data.SaveBlockResp{
		Block: block,
		Msg:   "",
		Code:  0,
	}, nil
}
