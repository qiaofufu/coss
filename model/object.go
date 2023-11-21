package model

type Object struct {
	ObjectID uint64 `json:"object_id" gorm:"autoIncrement"`
	Blocks   []byte `json:"blocks"`
	Data     []byte `json:"data" gorm:"-"`
}
