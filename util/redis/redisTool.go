package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

const RMQ string = "mqtest"

func producer() {
	redis_conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("hdiot"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer redis_conn.Close()

	rand.Seed(time.Now().UnixNano())

	var i = 1

	for {
		_, err = redis_conn.Do("rpush", RMQ, strconv.Itoa(i))
		if err != nil {
			fmt.Println("produce error")
			continue
		}
		fmt.Println("produce element:%d", i)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		i++
	}
}

func consumer() {
	redis_conn, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("hdiot"))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer redis_conn.Close()

	rand.Seed(time.Now().UnixNano())

	for {
		ele, err := redis.String(redis_conn.Do("lpop", RMQ))
		if err != nil {
			fmt.Println("no msg.sleep now")
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		} else {
			fmt.Println("cosume element:%s", ele)
		}
	}
}

func main() {
	list := os.Args
	if list[1] == "pro" {
		go producer()
	} else if list[1] == "con" {
		go consumer()
	}
	for {
		time.Sleep(time.Duration(10000) * time.Second)
	}
}
