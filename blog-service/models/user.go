package models

import "blog-service/utils"

type User struct {
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
}

func NewUser(fullName, userName, password, dateOfBirth string) *User {
	return &User{
		Id:          utils.GenerateUUID(),
		FullName:    fullName,
		UserName:    userName,
		Password:    password,
		DateOfBirth: dateOfBirth,
	}
}

func (u *User) getFullName() string {
	return u.FullName
}
