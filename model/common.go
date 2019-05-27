package model

import (
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

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

//初始化数据库连接
func InitDB(address, port, user, password, schema, charset string) {
	var err error
	param := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", user, password, address, port, schema)
	db, err = gorm.Open("mysql", param)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	if err != nil {
		panic("数据库连接失败")
	}

	/*
		检查表是否存在; 不存在新建
		struct对应创建表会在后面+s eq: type a struct 生成table as
		但是有些gorm指定特殊关键字对应规则会改变
		比如person 对应table people
	*/
	if err := db.AutoMigrate(
		new(Person),
	).Error; err != nil {
		panic(err)
	}
}
