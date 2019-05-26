package main

import (
	"flag"
	"gintest/cmd"
	"gintest/model"
)

func main() {
	flag.Parse()
	var config *model.TomlConfig = new(model.TomlConfig)
	if 0 != config.Init() {
		return
	}
	// go原生连接mysql 弃用
	// common.InitDB(config.MysqlConf.Host, config.MysqlConf.Port, config.MysqlConf.User, config.MysqlConf.Pass, config.MysqlConf.Schema, config.MysqlConf.Charset)

	//gorm
	model.InitDB(config.MysqlConf.Host, config.MysqlConf.Port, config.MysqlConf.User, config.MysqlConf.Pass, config.MysqlConf.Schema, config.MysqlConf.Charset)
	router := cmd.InitRouter()
	router.Run(config.ServerConf.Address)
}
