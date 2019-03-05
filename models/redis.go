package models

import (
	"fmt"
	//"log"
	//_"github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
)

func InitCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	return conn
}
