package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/jinzhu/gorm"
	"resource-backend/utils"
)

type User struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (bool, error) {
	var user User
	h := md5.New()
	h.Write([]byte(password))
	password = hex.EncodeToString(h.Sum(nil))
	err := db.Select("id").Where(User{Username:username, Password:password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetUserInfo(username, password string) (user User) {
	password = utils.EncodeMD5(password)
	db.Where(User{Username:username, Password:password}).First(&user)
	return
}

