package models

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func InitCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println(conn)
	return conn
}
