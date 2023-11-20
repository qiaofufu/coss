package model

import "coss/pkg/pg"

func Migrate() {
	err := pg.Client.AutoMigrate(
		&Meta{},
		&Object{},
	)
	if err != nil {
		panic(err)
	}
}
