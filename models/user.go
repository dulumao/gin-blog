package models

import (
	"errors"
	"gin-blog/pkg/utils"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"size:20;not null;unique;unique_index"`
	Password string `gorm:"not null"`
	Name     string `gorm:"size:20"`
	Phone    string `gorm:"size:20"`
	Model
}

// 设置表名
func (u User) TableName() string {
	return "blog_user"
}

func UserIsExistByUsername(username string) bool {
	var user User
	db.Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func CreateUser(username, password, name, phone string) (*User, error) {
	password, err := utils.Encrypt(password)
	if err != nil {
		return nil, err
	}
	user := &User{
		Username: username,
		Password: password,
		Name:     name,
		Phone:    phone,
	}
	i := db.Create(user).RowsAffected
	if i < 1 {
		return nil, errors.New("create user fail")
	}
	return user, nil
}

func CheckPassword(username, password string) *User {
	user := &User{}
	i := db.Where("username = ? ", username).First(user).RowsAffected
	if i != 1 {
		return nil
	}
	err := utils.CheckPassword(user.Password, password)
	if err == nil {
		return user
	}
	return nil
}
