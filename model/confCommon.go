package model

import "github.com/BurntSushi/toml"

type Server struct {
	Address string //项目监听端口
}
type Mysql struct {
	Host    string //mysql地址
	Port    string //端口
	User    string //用户名
	Pass    string //密码
	Schema  string
	Charset string //“utf8”
}
type TomlConfig struct {
	ServerConf *Server `toml:"server"`
	MysqlConf  *Mysql  `toml:"mysql"`
}

func (config *TomlConfig) Init() int {
	//读取配置文件
	_, err := toml.DecodeFile("config.toml", config)
	if err != nil {
		//log.info(err)
		return -1
	}
	return 0
}
