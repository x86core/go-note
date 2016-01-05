package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

var MAX_POOL_SIZE = 20 // pool size

var redisPoll chan redis.Conn // channel

func putRedis(conn redis.Conn) {
	if redisPoll == nil {
		redisPoll = make(chan redis.Conn, MAX_POOL_SIZE)
	}

	if len(redisPoll) >= MAX_POOL_SIZE {
		conn.Close()
		return
	}
	redisPoll <- conn
}

func InitRedis(network, address string) redis.Conn {

	if len(redisPoll) == 0 {
		redisPoll = make(chan redis.Conn, MAX_POOL_SIZE)

		go func() {
			for i := 0; i < MAX_POOL_SIZE/2; i++ {
				c, err := redis.Dial(network, address)

				if err != nil {
					panic(err)
				}

				putRedis(c)
			}
		}()
	}

	return <-redisPoll
}

func redisServer(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	c := InitRedis("tcp", "127.0.0.1:6379")
	dbkey := "redispoll:stack"
	if ok, err := redis.Bool(c.Do("LPUSH", dbkey, "pushval")); ok {

	} else {
		log.Print(err)
	}

	msg := fmt.Sprintf("use Time: %s", time.Now().Sub(startTime))

	io.WriteString(w, msg+"\n\n")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/", redisServer)
	http.ListenAndServe(":9527", nil)
}
