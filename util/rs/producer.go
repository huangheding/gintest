package rs

import "fmt"

func Produce(content string) (err error) {
	pConn := GetRedisConn()
	if pConn.Err() != nil {
		fmt.Println(pConn.Err().Error())
		return
	}
	defer pConn.Close()

	if _, err = pConn.Do("lpush", "redismq", content); err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
