package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type PersonalBill struct {
	Model
	Type     int     `json:"type"`
	Money    float64 `json:"money"`
	Year     string  `json:"year"`
	Month    string  `json:"month"`
	Day      string  `json:"day"`
	Category string  `json:"category"`
}

// 获取数据列表
func PersonalBillList(page int, pageSize int) (PersonalBills []PersonalBill, count int, err error) {
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&PersonalBills).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}
	return PersonalBills, count, nil
}

// 添加数据
func PersonalBillAdd(PersonalBill *PersonalBill) error {
	if err := db.Create(PersonalBill).Error; err != nil {
		return err
	}
	return nil
}

// 查看数据详情
func PersonalBillView(maps interface{}) (personalBill PersonalBill) {
	db.Where(maps).First(&personalBill)
	return
}

// 更新数据
func PersonalBillUpdate(personalBill *PersonalBill) (err error) {
	if err = db.Save(personalBill).Error; err != nil {
		return err
	}
	return nil
}

// 删除数据
func PersonalBillDelete(id int) (err error) {
	ret := db.Where("id = ?", id).Delete(&PersonalBill{})
	if ret.RowsAffected == 0 {
		return errors.New("账单不存在")
	}
	return nil
}
