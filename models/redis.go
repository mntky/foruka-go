package models

import (
	"github.com/gomodule/redigo/redis"
)

func NewPool(dbnum int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:	5,
		MaxActive:	10,
		Dial:		func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379", redis.DialDatabase(dbnum))
		},
	}
}
