package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type PersonalPhoto struct {
	Model
	Url         string `json:"url"`
	Description string `json:"description"`
}

func (v PersonalPhoto) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Url,
			validation.Required.Error("链接地址不得为空"),
			is.URL.Error("图片上传错误")),
	)
}

// 获取数据列表
func PersonalPhotoList(page int, pageSize int) (PersonalPhotos []PersonalPhoto, count int, err error) {
	db.Find(&PersonalPhotos).Count(&count)
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&PersonalPhotos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}
	return PersonalPhotos, count, nil
}

// 添加数据
func AddPersonalPhoto(PersonalPhoto *PersonalPhoto) error {
	// 数据验证
	err := PersonalPhoto.Validate()
	if err != nil {
		return err
	}
	if err := db.Create(PersonalPhoto).Error; err != nil {
		return err
	}
	return nil
}

// 删除数据
func DeletePersonalPhoto(id int) (err error) {
	ret := db.Where("id = ?", id).Delete(&PersonalPhoto{})
	if ret.RowsAffected == 0 {
		return errors.New("视频不存在")
	}
	return nil
}
