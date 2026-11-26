package handler

import (
	"github.com/Qoiuy/go-base/backend/api/config"
	"github.com/Qoiuy/go-base/backend/flog"
	"github.com/Qoiuy/go-base/backend/mysql"
)

func InitDB() *mysql.Repository {
	// Init database
	datastoreCfg := &mysql.Config{
		HostAddress: config.Opts.MysqlConfig.MysqlAddress,
		Username:    config.Opts.MysqlConfig.MysqlUser,
		Password:    config.Opts.MysqlConfig.MysqlPassword,
		Database:    config.Opts.MysqlConfig.MysqlDatabase,
		//EnableLogging: false,
		EnableLogging: config.Opts.Debug,
	}
	orm, err := mysql.CreateRepository(datastoreCfg)
	if err != nil {
		flog.Infof("p:%+s", datastoreCfg.Password)
		flog.Fatal(err)
	}
	return orm
}
