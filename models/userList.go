package models

import "github.com/google/uuid"

type UserList struct {
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	CreaterId uuid.UUID `json:"createrId"`
}

type UserData struct {
	FullName string `json:"fullname"`
	Email    string `gorm:"primaryKey" json:"email"`
}
