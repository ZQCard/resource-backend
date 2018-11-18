package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"resource-backend/utils"
)

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

func (u User)Validate() error {
	return validation.ValidateStruct(&u,
		// 名称不得为空,且大小为1-30字
		validation.Field(
			&u.Username,
			validation.Required.Error("名称不得为空"),
			validation.Length(1, 50).Error("名称为1-50字")),
		// 头像不得为空,且为url地址
		validation.Field(&u.Avatar,
			validation.Required.Error("头像不得为空"),
			is.URL.Error("头像格式错误")),
	)
}

func GetUserByMaps(maps interface{}) (user User, err error) {
	err = db.Where(maps).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func UserList(page int, pageSize int) (users []User, count int, err error)   {
	err = db.Unscoped().Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Count(&count).Error
	if err != nil{
		return
	}
	return
}

func AddUser(user *User) error {
	err := user.Validate()
	if err != nil {
		return err
	}
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UserUpdate(user *User) (err error) {
	// 数据验证
	err = user.Validate()
	if err != nil {
		return err
	}
	if err = db.Debug().Save(user).Error; err != nil {
		return err
	}
	return nil
}

func UserDelete(id int) error {
	if err := db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func UserRecover(id int) {
	db.Model(&User{}).Where("id = ?", id).Unscoped().Update("deleted_at", nil)
}
