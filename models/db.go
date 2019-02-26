package models //Model

import (
	"github.com/go-gorm/xorm"
	_"github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

type User struct {
	ID			int		`json:"id"		xorm:"'id'"`
	Username	string	`json:"user"	xorm:"'user'"`	
}

type UserRepository struct {
}

//connect mysql
func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "mntky:123qwEcc@/forukago")
	if err != nil {
		panic(err)
	}
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
	return UserRepostitory{}
}

func (m UserRepostiroy) GetByID(id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}
	return nil
}
