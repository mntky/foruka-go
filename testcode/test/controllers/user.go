package controllers

import (
	"github.com/mntky/foruka-go/models"
)

type User struct {
}

func NewUser() User {
	return  User{}
}

func (c User) Get(n int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetByID(n)
	return user
}


