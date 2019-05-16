package main

import (
	"flag"
	"gintest/cmd"
	"gintest/common"
	"gintest/model"
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
