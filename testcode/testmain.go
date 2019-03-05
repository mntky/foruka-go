package main

import (
	"net/http"
	"log"
	//"github.com/gin-gonic/contrib/sessions"
	//"html/template"
	"github.com/mntky/foruka-go/testmodels"
)


func main() {
	testmodels.InitCache()

	http.HandleFunc("/signin", testmodels.Signin)
	http.HandleFunc("/welcome", testmodels.Welcome)
	http.HandleFunc("/test", testmodels.Test)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
