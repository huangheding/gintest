package main

import (
	"flag"
	"gin_test/cmd"
	"gin_test/common"
	"gin_test/model"
)

func main() {
	flag.Parse()
	var config *model.TomlConfig = new(model.TomlConfig)
	if 0 != config.Init() {
		return
	}
	common.InitDB(config.MysqlConf.Host, config.MysqlConf.Port, config.MysqlConf.User, config.MysqlConf.Pass, config.MysqlConf.Schema, config.MysqlConf.Charset)

	router := cmd.InitRouter()
	router.Run(config.ServerConf.Address)
}
