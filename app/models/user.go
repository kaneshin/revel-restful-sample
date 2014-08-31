package models

import (
	"errors"
	"net/url"
	"revel-boilerplate/app/entities"
)

type User struct {
	Model
}

var sharedInstance *User

func GetUserInstance() *User {
	if sharedInstance == nil {
		sharedInstance = NewUserInstance()
	}
	return sharedInstance
}

func NewUserInstance() *User {
	instance := &User{}
	instance.Init()
	return instance
}

func (self *User) Create(params url.Values) error {
	name := params.Get("name")
	if len(name) == 0 {
		return errors.New("name value is empty")
	}
	user := entities.NewUser(name)
	return self.Dbm.Insert(&user)
}

func (self *User) Get(id int64) (*entities.User, error) {
	user := entities.User{}
	err := self.Dbm.SelectOne(&user, "select * from users where id=?", id)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (self *User) Update(id int64, params url.Values) error {
	user, err := self.Get(id)
	if err != nil {
		return err
	}
	name := params.Get("name")
	if len(name) == 0 {
		return errors.New("name value is empty")
	}
	user.Name = name
	_, err = self.Dbm.Update(user)
	return err
}

func (self *User) Delete(id int64) error {
	user, err := self.Get(id)
	if err != nil {
		return err
	}
	_, err = self.Dbm.Delete(user)
	return err
}
