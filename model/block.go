package model

type Block struct {
	BlockID uint64 `json:"block_id" gorm:"autoIncrement;"`
	Data    []byte `json:"data"`
}
