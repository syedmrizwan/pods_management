package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func GetPool() *redis.Pool {
	return pool
}
