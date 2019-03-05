package models

import (
	"github.com/gomodule/redigo/redis"
)

func InitCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	cache := conn
	return cache
}
