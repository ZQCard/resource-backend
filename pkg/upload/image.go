package upload

import (
	"log"
	"mime/multipart"
	"resource-backend/pkg/config"
	"resource-backend/pkg/file"
	"strings"
)

type Image struct {
	Base
	ImageAllowExts []string
	ImageSavePath  string
	ImageMaxSize   int
}

var image = &Image{}

func init() {
	image.ImageAllowExts = config.AppSetting.ImageAllowExts
	image.ImageSavePath = config.AppSetting.ImageSavePath
	image.ImageMaxSize = config.AppSetting.ImageMaxSize
}

func (i *Image)GetPath() string {
	return image.ImageSavePath
}

func (i *Image)GetFullUrl(name string)  string {
	return i.GetFullBaseUrl() + "/" + i.GetPath() + name
}


func (i *Image)CheckExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range image.ImageAllowExts {
		if strings.ToLower(allowExt) == strings.ToLower(ext) {
			return true
		}
	}
	return false
}

func (i *Image)CheckSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return size <= image.ImageMaxSize
}
