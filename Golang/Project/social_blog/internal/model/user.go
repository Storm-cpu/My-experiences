package model

import "time"

type User struct {
	ID          int        `json:"id" gorm:"primary_key"`
	FirstName   string     `json:"first_name" gorm:"type:varchar(255)"`
	LastName    string     `json:"last_name" gorm:"type:varchar(255)"`
	PhoneNumber string     `json:"phone_number,omitempty" gorm:"type:varchar(20)"`
	Email       string     `json:"email" gorm:"type:varchar(255)"`
	LastLogin   *time.Time `json:"last_login,omitempty"`
	UserName    string     `json:"username" gorm:"type:varchar(255);unique_index;not null"`
	Password    string     `json:"-" gorm:"type:varchar(255);not null"`
	Blocked     bool       `json:"blocked" gorm:"not null;default:0"`
	Base
}
