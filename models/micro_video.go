package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type MicroVideo struct {
	Model
	Url  string `json:"url"`
	View int `json:"view"`
}

func (v MicroVideo) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Url,
			validation.Required.Error("链接地址不得为空"),
			is.URL.Error("视频上传错误")),
	)
}

// 获取数据列表
func MicroVideoList(page int, pageSize int) (microVideos []MicroVideo, count int, err error) {
	db.Find(&microVideos).Count(&count)
	err = db.Select("url,view").Offset((page - 1) * pageSize).Limit(pageSize).Order("id DESC").Find(&microVideos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}
	return microVideos, count, nil
}

// 添加数据
func AddMicroVideo(microVideo *MicroVideo) error {
	// 数据验证
	err := microVideo.Validate()
	if err != nil {
		return err
	}
	if err := db.Create(microVideo).Error; err != nil {
		return err
	}
	return nil
}

func FindMicroVideoByUrl(url string) bool {
	var video MicroVideo
	db.Model(MicroVideo{}).Select("id").Where("url = ?", url).First(&video)
	if video.ID > 0 {
		return true
	}
	return false
}

// 查看数据详情
func ViewMicroVideo(url string) (err error) {
	res := db.Model(MicroVideo{}).Where("url = ?", url).UpdateColumn("view", gorm.Expr("view + ?", 1))
	if res.RowsAffected == 0 {
		return errors.New("视频不存在")
	}
	return nil
}

// 删除数据
func DeleteMicroVideo(url string) (err error) {
	ret := db.Where("url = ?", url).Delete(&MicroVideo{})
	if ret.RowsAffected == 0 {
		return errors.New("视频不存在")
	}
	return nil
}
