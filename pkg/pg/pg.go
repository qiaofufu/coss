package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func MustInitPG(host, user, passwd, dbname, port string) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host,
		user,
		passwd,
		dbname,
		port,
	)
	db, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		panic(err)
	}
	Client = db
}
