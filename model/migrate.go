package model

import "coss/pkg/pg"

func Migrate() {
	err := pg.Client.AutoMigrate(
		&Meta{},
		&Object{},
		&Block{},
	)
	if err != nil {
		panic(err)
	}
}
