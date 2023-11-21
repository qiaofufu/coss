package controller

import (
	"coss/model"
	"coss/pb/data"
	"gorm.io/gorm"
)

type DataController interface {
	SaveBlock(data []byte) (*data.Block, error)
	GetBlock(blockID uint64, tableName string, offset int64, limit int64) ([]byte, error)
}

func NewDataController(db *gorm.DB) DataController {
	return &dataCtrl{
		db: db,
	}
}

type dataCtrl struct {
	db *gorm.DB
}

func (d *dataCtrl) SaveBlock(dat []byte) (*data.Block, error) {
	var (
		block   = new(model.Block)
		blockPB = new(data.Block)
		err     error
	)
	block.Data = dat
	err = d.db.Create(block).Error
	if err != nil {
		return nil, err
	}
	blockPB.BlockID = block.BlockID
	blockPB.Size = int64(len(dat))
	blockPB.TableName = "block"

	return blockPB, nil
}

func (d *dataCtrl) GetBlock(blockID uint64, tableName string, offset int64, limit int64) ([]byte, error) {
	var (
		tx    = d.db.Model(model.Block{})
		block = new(model.Block)
		err   error
	)
	tx.Where("block_id = ?", blockID)
	err = tx.Take(block).Error
	if err != nil {
		return nil, err
	}
	return block.Data[offset : offset+limit], nil
}
