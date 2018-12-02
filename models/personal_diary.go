package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type PersonalDiary struct {
	Model
	Title string `json:"title"`
	Content string `json:"content"`
	Secret string `json:"secret"`
}

// 获取数据列表
func PersonalDiaryList(page int, pageSize int) (PersonalDiarys []PersonalDiary, count int, err error) {
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&PersonalDiarys).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}
	return PersonalDiarys, count, nil
}

// 添加数据
func PersonalDiaryAdd(PersonalDiary *PersonalDiary) error {
	if err := db.Create(PersonalDiary).Error; err != nil {
		return err
	}
	return nil
}

// 查看数据详情
func PersonalDiaryView(maps map[string]interface{}) (PersonalDiary PersonalDiary) {
	db.Where(maps).First(&PersonalDiary)
	return
}

// 更新数据
func PersonalDiaryUpdate(PersonalDiary *PersonalDiary) (err error) {
	if err = db.Save(PersonalDiary).Error; err != nil {
		return err
	}
	return nil
}

// 删除数据
func PersonalDiaryDelete(id int) (err error) {
	ret := db.Where("id = ?", id).Delete(&PersonalDiary{})
	if ret.RowsAffected == 0 {
		return errors.New("日记不存在")
	}
	return nil
}
