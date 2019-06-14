package rs

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func Custom(address, key string) string {
	c := GetRedisConn()

	defer c.Close()

	ele, err := redis.Strings(c.Do("brpop", key, 0))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("cosume element:%s", ele)
	var data string
	if len(ele) > 1 {
		data = ele[1]
		return data
	}
	return data
}
