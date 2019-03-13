package main

import (
	"golang.org/x/crypto/bcrypt"
)

func main() {
	    hash, err := passwordHash("test1234")
    if err != nil {
        panic(err)
    }
    println(hash)
}

// パスワードハッシュを作る
func passwordHash(pw string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), err
}
