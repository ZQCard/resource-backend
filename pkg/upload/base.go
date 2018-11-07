package upload

import (
	"fmt"
	"os"
	"path"
	"resource-backend/pkg/config"
	"resource-backend/pkg/file"
	"resource-backend/utils"
	"strconv"
	"strings"
)

type Base struct {
	FullBaseUrl string
}

var base = &Base{}

func init()  {
	base.FullBaseUrl = config.AppSetting.BaseUrl + ":" + strconv.Itoa(config.ServerSettings.HTTPPort)
}

func (b *Base)GetFullBaseUrl() string {
	return base.FullBaseUrl
}

func (b *Base)GetName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)
	return fileName + ext
}


func (b *Base)CheckExist(src string) bool {
	dir, err := os.Getwd()
	if err != nil {
		return false
	}
	if !file.CheckExist(dir + "/" + src) {
		return false
	}
	return true
}

func (b *Base)MakePath(src string) error {
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





