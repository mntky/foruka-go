package main

import(
	"fmt"
	"log"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


type User struct {
	ID		int		`db:"id"`
	Name	string	`db:"name"`
	Pass	string	`db:"password"`
}

type Userlist []User

func main() {
	var userlist Userlist

	db, err := sqlx.Open("mysql", "mntky:123qwEcc@/forukago")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Queryx("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}

	var user User
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil{
			log.Fatal(err)
		}
		userlist = append(userlist, user)
	}

	fmt.Println(userlist)
}
