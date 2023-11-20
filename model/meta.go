package model

type Meta struct {
	Bucket   string `gorm:"not null;"`
	Key      string `gorm:"not null"`
	ObjectID uint64 `gorm:"not null"`
}
