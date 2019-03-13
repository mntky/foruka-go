package main

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"unsafe"
)

func create(password string) string {
	b := *(*[]byte)(unsafe.Pointer(&password))
	salt := base64.StdEncoding.EncodeToString(b)
	fmt.Println(salt)
	converted, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 16)
	return hex.EncodeToString(converted[:])
}

func main() {
	text := "hello"
	ans := create(text)

	fmt.Println(ans)

}
