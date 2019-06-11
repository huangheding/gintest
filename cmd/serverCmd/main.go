package main

import (
	"flag"
	"fmt"
	"gintest/cmd"
	"gintest/model"
	"gintest/util/ws"
)

func init() {
	fmt.Println("init")
}

func main() {
	flag.Parse()
	var config *model.TomlConfig = new(model.TomlConfig)
	if 0 != config.Init() {
		return
	}
	//gorm
	model.InitDB(config.MysqlConf.Host, config.MysqlConf.Port, config.MysqlConf.User, config.MysqlConf.Pass, config.MysqlConf.Schema, config.MysqlConf.Charset)
	//ws
	ws.InitWs()
	router := cmd.InitRouter()
	router.Run(config.ServerConf.Address)
}
