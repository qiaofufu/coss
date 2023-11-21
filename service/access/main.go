package main

import (
	"coss/model"
	"coss/opt"
	"coss/pkg/pg"
	"coss/service/access/internal/api"
)

func main() {
	opt.MustInitConfig()
	pg.MustInitPG(opt.Cfg.Pg.Host, opt.Cfg.Pg.User, opt.Cfg.Pg.Passwd, opt.Cfg.Pg.DBName, opt.Cfg.Pg.Port)
	model.Migrate()
	api.MustStartAPI(opt.Cfg.Access.Port)
}
