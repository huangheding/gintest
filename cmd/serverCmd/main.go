package main

import (
	"flag"
	"gintest/cmd"
	"gintest/model"
	"gintest/util/rs"
	"gintest/util/ws"
)

func init() {

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
	//redis
	rs.InitRedis(config.RedisConf.Adress)

	router := cmd.InitRouter()
	router.Run(config.ServerConf.Address)
}
