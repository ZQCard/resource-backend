package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"resource-backend/pkg/config"
	"resource-backend/pkg/file"
	"resource-backend/utils"
	"strconv"
	"strings"
)

func GetFullBaseUrl() string {
	return config.AppSetting.BaseUrl + ":" + strconv.Itoa(config.ServerSettings.HTTPPort)
}

func GetFullUrl(name string)  string {
	return GetFullBaseUrl() + "/" + GetPath() + name
}

func GetName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}

func GetPath() string {
	return config.AppSetting.ImageSavePath
}

func CheckExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range config.AppSetting.ImageAllowExts{
		if strings.ToLower(allowExt) == strings.ToLower(ext) {
			return true
		}
	}
	return  false
}

func CheckSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return size <= config.AppSetting.ImageMaxSize
}

func CheckExist(src string) bool {
	dir, err := os.Getwd()
	if err != nil {
		return false
	}
	if !file.CheckExist(dir + "/" + src) {
		return false
	}
	return true
}

func MakePath(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}
	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}
	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}