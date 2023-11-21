package controller

import (
	"coss/model"
	"coss/opt"
	"coss/pb/data"
	"coss/pb/meta"
	"coss/pkg/test"
	"coss/pkg/tools"
	"fmt"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
	"mime/multipart"
)

const (
	MaxSplitSize int64 = 1024 * 1024
	Timeout            = 60 * 10
)

type AccessControlller interface {
	CreateObject(bucketName, objectName string, file *multipart.FileHeader) error
	GetObject(bucketName, objectName string) (*model.Object, error)
	GetObjectWithData(bucketName, objectName string) (*model.Object, error)
}

func NewAccessController(db *gorm.DB) AccessControlller {
	dataAddr := fmt.Sprintf("%s:%s", opt.Cfg.Data.Host, opt.Cfg.Data.Port)
	dataRPC := test.GetClient(dataAddr, data.NewDataClient)
	metaAddr := fmt.Sprintf("%s:%s", opt.Cfg.Meta.Host, opt.Cfg.Meta.Port)
	metaRPC := test.GetClient(metaAddr, meta.NewMetaClient)
	return &accessCtrl{
		db:      db,
		dataRPC: dataRPC,
		metaRPC: metaRPC,
	}
}

type accessCtrl struct {
	db      *gorm.DB
	metaRPC meta.MetaClient
	dataRPC data.DataClient
}

func (a *accessCtrl) GetObject(bucketName, objectName string) (*model.Object, error) {
	var (
		object = new(model.Object)
		tx     = a.db.Model(object)
	)
	resp, err := a.metaRPC.GetMeta(tools.Timeout(3), &meta.GetMetaReq{
		Bucket: bucketName,
		Key:    objectName,
	})
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("get meta failed, err: %v", err)
	}
	tx.Where("object_id = ?", resp.ObjectID)
	err = tx.Take(object).Error
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (a *accessCtrl) GetObjectWithData(bucketName, objectName string) (*model.Object, error) {
	var (
		object *model.Object
		err    error
		blocks = new(data.DataBlocks)
	)
	object, err = a.GetObject(bucketName, objectName)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(object.Blocks, blocks)
	if err != nil {
		return nil, err
	}

	for _, block := range blocks.GetBlocks() {
		resp, err := a.dataRPC.DownloadBlock(tools.Timeout(), &data.DownloadBlockReq{
			BucketID:  block.BlockID,
			Offset:    0,
			Limit:     block.Size,
			TableName: block.TableName,
		})
		if err != nil {
			return nil, err
		}
		if resp.Code != 0 {
			return nil, fmt.Errorf("download block failed, err: %v", err)
		}
		object.Data = append(object.Data, resp.GetData()...)
	}
	return object, nil
}
