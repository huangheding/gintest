package rs

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gomodule/redigo/redis"
)

func Custom(address, key string) string {
	c, err := redis.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer c.Close()
	for {
		ele, err := redis.String(c.Do("lpop", key))
		if err != nil {
			// fmt.Println("no msg.sleep now")
			time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
		} else {
			fmt.Println("cosume element:%s", ele)
			return ele
		}
	}
}
