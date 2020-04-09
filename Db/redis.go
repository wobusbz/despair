package Db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	Pool = new(redis.Pool)
)

const (
	RedisPass = ""
	NOTFOUND  = "Unknown"
)

// RedisUrl  = "redis://localhost:6379"
func InitRedis(redisUrl string) {
	once.Do(func() {
		Pool = &redis.Pool{
			MaxIdle:     1217,  // 最大空闲连接数
			MaxActive:   0,     // 最大连接数
			IdleTimeout: 0,     // 最大空闲连接等待时间
			Wait:        false, // 为true在Get是会受最大连接数限制
			Dial: func() (conn redis.Conn, err error) {
				conn, err = redis.DialURL(redisUrl)
				if err != nil {
					return nil, fmt.Errorf("redis connection error: %s", err)
				}

				//if _, err := conn.Do("AUTH", RedisPass); err != nil { //
				//	return nil, fmt.Errorf("redis auth password error: %s", err)
				//}
				return
			},
		}
	})
}
