package models

import (
	"encoding/hex"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"unsafe"
)

func Create(password, potato string) (string, string){
	b := *(*[]byte)(unsafe.Pointer(&potato))
	salt := base64.StdEncoding.EncodeToString(b)
	converted, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 16)
	return hex.EncodeToString(converted[:]), salt
}
