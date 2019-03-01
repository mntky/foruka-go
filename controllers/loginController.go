package controllers

import (
	"github.com/mntky/foruka-go/models"
)

type POSTdata struct {
	Username	string	`json:"username"`
	Password	string	`json:"password"`
}

//Default Login Page
func Login(g *gin.Context) {
	c.HTML(200, "login2.tmpl", nil)
}

//receive Post data from login2.tmpl
//processing Posted JSON data(username, password) and Send to SigninFunc
func Auth(g *gin.Context) {
	var data POSTdata
	g.BindJSON($data)
	fmt.Printf("ok %s", data)
	err := Signin(g, data.Username, data.Password)
	if err != nil {
		fmt.Println(err)
	}
}

//Login Process 
//use username and pass from AuthFunc 
func Signin(g *gin.Context, uname, pass string) err string{
	//Connect to Redis Server
	models.InitCache()

	//check password
	getpass, check := user[uname]
	if !ok || getpass != pass {
		return "Username or Password is Not Correct"
	}

	//Create token of cookie, and it in sessionToken
	u, err := uuid.NewV4()
	sessionToken := u.String

	//Set SessionToken to Redis 
	_, err = cache.Do("SETEX", sessionToken, "120", uname)
	if err != nil {
		return "cant set cookie token"
	}

	//Set Cookie to User
	http.SetCookie(g.Writer, &http.Cookie{
		Name:		"session_token",
		Value:		sessionToken,
		Expires:	time.Now().Add(120 * time.Second),
	})

	//Redirect to LoginSuccessPage
	g.Redirect(http.StatusMovedPermanently, "http://192.168.11.100/welcome")

