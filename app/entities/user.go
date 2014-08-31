package entities

import (
	"time"
)

type User struct {
	Id      int64  `db:"id" json:"id"`
	Name    string `db:"name" json:"name"`
	Created int64  `db:"created" json:"created"`
}

func NewUser(name string) User {
	return User{
		Name:    name,
		Created: time.Now().UnixNano(),
	}
}
