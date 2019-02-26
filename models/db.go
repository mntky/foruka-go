package models //Model

import (
	//"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
)


type User struct {
	ID			int		`json:"id"		xorm:"'id'"`
	Username	string	`json:"user"	xorm:"'user'"`	
}

type UserRepository struct {
}

//connect mysql
func NewMySQL(dsn string) (*sqlx.DB, error) {
	DB, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
}


//newuser
func NewUser(id int, username string) User {
	return User {
		ID:			id,
		Username: 	username,
	}
}

//NewUserRepository
func NewUserRepostiroy() UserRepository {
	return UserRepository{}
}

func (m UserRepository) GetByID(id int) *User {
	var user = User{ID: id}
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
    if err != nil {
        fmt.Println(err)
    }
	//err := db.Get(&user)
	if err {
		return &user
	}
	return nil
}
