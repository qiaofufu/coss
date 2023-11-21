package controller

import (
	"coss/model"
	"coss/pb/data"
	"coss/pb/meta"
	"coss/pkg/tools"
	"fmt"
	"google.golang.org/protobuf/proto"
	"mime/multipart"
)

func (a *accessCtrl) CreateObject(bucketName, objectName string, fileHeader *multipart.FileHeader) error {
	var (
		err    error
		file   multipart.File
		size   int64
		object = new(model.Object)
		blocks = new(data.DataBlocks)
	)

	resp, err := a.metaRPC.Exist(tools.Timeout(3), &meta.ExistReq{
		Bucket: bucketName,
		Key:    objectName,
	})
	if err != nil {
		return err
	}
	if resp.GetResult() {
		return fmt.Errorf("object already exist")
	}

	file, err = fileHeader.Open()
	if err != nil {
		return fmt.Errorf("failed open file, err: %v", err)
	}

	size = fileHeader.Size

	for i := int64(0); size-i*MaxSplitSize >= 0; i++ {
		buf := make([]byte, MaxSplitSize)
		n, err := file.Read(buf)
		if err != nil {
			return fmt.Errorf("failed read file part, part: %d, err: %v", i, err)
		}
		buf = buf[:n]
		resp, err := a.dataRPC.SaveBlock(tools.Timeout(Timeout), &data.SaveBlockReq{Data: buf})
		if resp.Code != 0 {
			return fmt.Errorf("data rpc error: code: %d, msg: %s", resp.Code, resp.Msg)
		}
		blocks.Blocks = append(blocks.Blocks, resp.Block)
	}
	bytes, err := proto.Marshal(blocks)
	if err != nil {
		return err
	}
	object.Blocks = bytes

	err = a.db.Create(object).Error
	if err != nil {
		return err
	}

	res, err := a.metaRPC.CreateMeta(tools.Timeout(60), &meta.CreateMetaReq{
		Bucket:   bucketName,
		Key:      objectName,
		ObjectID: object.ObjectID,
	})
	if err != nil {
		return err
	}
	if res.Code != 0 {
		return fmt.Errorf("meta rpc error: code: %d, msg: %s", res.Code, res.Msg)
	}

	return nil
}
