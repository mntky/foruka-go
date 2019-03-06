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



/*
func InitCache() redis.Conn{
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println(conn)
	return conn
}
*/

