package rs

import "fmt"

func Produce(szBytes []byte) (err error) {
	pConn := GetRedisConn()
	if pConn.Err() != nil {
		fmt.Println(pConn.Err().Error())
		return
	}
	defer pConn.Close()

	if _, err = pConn.Do("lpush", "redismq", szBytes); err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}
