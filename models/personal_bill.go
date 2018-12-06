package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

const PersonalBillExpand  = 0
const PersonalBillIncome  = 1

type PersonalBill struct {
	Model
	Type     int     `json:"type"`
	Money    float64 `json:"money"`
	Year     string  `json:"year"`
	Month    string  `json:"month"`
	Day      string  `json:"day"`
	Category string  `json:"category"`
}

func PersonalBillSummary() map[int]float64 {
	rows, err := db.Model(&PersonalBill{}).Select("SUM(money) as value, type as kind").Group("type").Rows()
	summary := make(map[int]float64)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var money float64
			var kind int
			rows.Scan(&money, &kind)
			summary[kind] = money
		}
	}
	return summary
}

func PersonalBillSummaryByCategory(maps interface{}) interface{} {
	var data []map[string]interface{}
	rows, err := db.Model(&PersonalBill{}).Where(maps).Select("SUM(money) as value, category as name").Group("category").Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			temp := make(map[string]interface{})
			var money float64
			var name string
			rows.Scan(&money, &name)
			temp["name"] = name
			temp["value"] = money
			data = append(data, temp)
		}
	}
	return data
}

func PersonalBillSummaryByType(maps interface{}) interface{} {
	var data []map[string]interface{}
	rows, err := db.Model(&PersonalBill{}).Where(maps).Select("SUM(money) as value, type as kind").Group("type").Rows()
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			temp := make(map[string]interface{})
			var money float64
			var kind int
			rows.Scan(&money, &kind)
			if kind == PersonalBillExpand{
				temp["name"] = "支出"
			}else {
				temp["name"] = "收入"
			}
			temp["value"] = money
			data = append(data, temp)
		}
	}
	return data
}

func PersonalBillSummaryByYear(maps interface{}) (map[string]float64, map[string]float64) {
	var moneys []PersonalBill
	db.Model(&PersonalBill{}).Where(maps).Select("money, type, month").Order("month").Find(&moneys)
	moneysIncome := make(map[string]float64)
	moneysExpand := make(map[string]float64)
	for _, v := range moneys{
		if v.Type == PersonalBillExpand {
			moneysExpand[v.Month] += v.Money
		}else {
			moneysIncome[v.Month] += v.Money
		}
	}
	return moneysExpand, moneysIncome
}

// 获取数据列表
func PersonalBillList(page int, pageSize int) (PersonalBills []PersonalBill, count int, err error) {
	db.Find(&PersonalBills).Count(&count)
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&PersonalBills).Error
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
