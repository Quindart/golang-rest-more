package models

import "blog-service/utils"

type User struct {
	Id          string `json:"id"`
	FullName    string `json:"full_name"`
	UserName    string `json:"user_name"`
	Password    string `json:"-"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func NewUser(fullName, userName, password, dateOfBirth, phoneNumber, email string) *User {
	return &User{
		Id:          utils.GenerateUUID(),
		FullName:    fullName,
		UserName:    userName,
		Password:    password,
		DateOfBirth: dateOfBirth,
		PhoneNumber: phoneNumber,
		Email:       email,
	}
}

func (u *User) getFullName() string {
	return u.FullName
}
