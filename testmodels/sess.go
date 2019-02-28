package testmodels

import (
	//"log"
	"net/http"
	"github.com/gomodule/redigo/redis"
	"fmt"
	"encoding/json"
	"github.com/satori/go.uuid"
	"time"
)

var cache redis.Conn

var user = map[string]string {
	"user1":	"123qwEcc",
	"uesr2":	"123qwEcc@",
}

type Credentials struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

func InitCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	cache = conn
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var c Credentials
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := user[c.Username]

	if !ok || expectedPassword != c.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	u, err := uuid.NewV4()
	sessionToken := u.String()

	_, err = cache.Do("SETEX", sessionToken, "120", c.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:		"session_token",
		Value:		sessionToken,
		Expires:	time.Now().Add(120 * time.Second),
	})
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if response == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("welcome %s ! ", response)))
}


