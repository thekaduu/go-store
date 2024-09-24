package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	StoreID  uint `gorm:"index"`
	Store    Store
	Username string `gorm:"unique"`
	Password string
	Role     string
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return
	}

	u.Password = string(bytes)
	return
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) EncryptedPassword() string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		return ""
	}

	return string(bytes)
}
