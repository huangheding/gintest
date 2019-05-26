package common

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

var dbHost, dbPort, dbSchema, dbName, dbPass = "", "", "", "", ""

//初始化DB
func InitDB(address, port, user, password, schema, charset string) {
	var err error
	param := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local", user, password, address, port, schema, charset)
	dbInstance, err = sql.Open("mysql", param)
	//错误检查
	if err != nil {
		log.Fatal(err.Error())
	}
	//推迟数据库连接的关闭
	// defer Db.Close()

	//
	err = dbInstance.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	dbHost = strings.TrimSpace(address)
	dbPort = strings.TrimSpace(port)
	dbSchema = strings.TrimSpace(schema)
	dbName = strings.TrimSpace(user)
	dbPass = strings.TrimSpace(password)
}

//获取db实例 如果为nil就刷新连接
func GetDb() (*sql.DB, error) {
	if nil == dbInstance {
		param := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", dbName, dbPass, dbHost, dbPort, dbSchema)
		dbStr, err := sql.Open("mysql", param)
		if err != nil {
			return nil, err
		}
		//设置
		dbStr.SetMaxOpenConns(1000)
		dbStr.SetMaxIdleConns(10)
		dbStr.SetConnMaxLifetime(time.Hour * 5)
		dbInstance = dbStr
	}

	err := dbInstance.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return dbInstance, nil
}

// 封装查询接口
func QueryTableData(sqlStr string) (*sql.Rows, error) {
	var db *sql.DB
	var err error
	if db, err = GetDb(); err != nil {
		return nil, err
	}
	rows, err := db.Query(sqlStr)
	if err != nil {
		return rows, err
	}

	return rows, err
}
