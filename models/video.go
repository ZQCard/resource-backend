package models

import (
	"github.com/jinzhu/gorm"
)

type Videos struct {
	Model
	Type int8 `json:"type"`
	Name string `json:"name"`
	Href string `json:"href"`
}

// 获取记录总数
func GetVideosTotalCount(maps interface{}) (count int) {
	db.Model(&Videos{}).Where(maps).Count(&count)
	return
}

// 获取数据列表
func GetVideosList(page int, pageSize int, maps interface{}) ([]Videos, error)  {
	var videos []Videos

	err := db.Where(maps).Offset(page - 1).Limit(pageSize).Find(&videos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return videos, nil
}