package controller

import (
	"coss/model"
	"fmt"
	"gorm.io/gorm"
)

type MetaController interface {
	CreateMeta(bucket, key string, objectID uint64) error
	Exist(bucket, key string) bool
	GetMeta(bucket, key string) (*model.Meta, error)
}

func NewMetaController(db *gorm.DB) MetaController {
	ctrl := &metaCtrl{
		db,
	}
	return ctrl
}

type metaCtrl struct {
	db *gorm.DB
}

func (m *metaCtrl) CreateMeta(bucket, key string, objectID uint64) error {
	meta := model.Meta{
		Bucket:   bucket,
		Key:      key,
		ObjectID: objectID,
	}
	if m.Exist(bucket, key) {
		return fmt.Errorf("meta info already exist")
	}
	if err := m.db.Create(meta).Error; err != nil {
		return err
	}
	return nil
}

func (m *metaCtrl) Exist(bucket, key string) bool {
	var (
		tx  = m.db.Model(&model.Meta{})
		cnt int64
	)
	tx.Where("bucket = ? AND key = ?", bucket, key).Count(&cnt)
	return cnt > 0
}

func (m *metaCtrl) GetMeta(bucket, key string) (*model.Meta, error) {
	var (
		meta = new(model.Meta)
		tx   = m.db.Model(meta)
		err  error
	)
	tx.Where("bucket = ? and key = ?", bucket, key)
	err = tx.Take(meta).Error
	return meta, err
}
