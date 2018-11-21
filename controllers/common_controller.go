package controllers

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io/ioutil"
	"net/http"
	"resource-backend/pkg/config"
	"resource-backend/pkg/logging"
	"resource-backend/pkg/upload"
)

var (
	accessKey   = config.GetConfigParam("QINIU", "QINIU_ACCESS_KEY")
	secretKey   = config.GetConfigParam("QINIU", "QINIU_SECRET_KEY")
	bucket      = config.GetConfigParam("QINIU", "QINIU_BUCKET")
	url         = config.GetConfigParam("QINIU", "QINIU_CDN_DOMAIN")
	callbackUrl = config.GetConfigParam("QINIU", "QINIU_CALLBACK_URL")
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"asdas": "sadsa",
	})
}

func Upload(c *gin.Context) {
	fType := c.PostForm("type")
	switch fType {
	case "image":
		processImage(c)
	default:
		c.JSON(http.StatusOK, gin.H{
			"message": "请求参数错误",
		})
	}
	return
}

func processImage(c *gin.Context) {
	var code = http.StatusOK
	data := make(map[string]interface{})
	file, fHeader, err := c.Request.FormFile("file")
	if err != nil || fHeader == nil {
		data["message"] = "文件上传错误:" + err.Error()
		c.JSON(code, data)
		logging.Error(err.Error())
		return
	}
	uploadType := &upload.Image{}

	fName := uploadType.GetName(fHeader.Filename)
	savePath := uploadType.GetPath()

	src := savePath + fName

	if !uploadType.CheckExt(fName) || !uploadType.CheckSize(file) {
		data["message"] = "文件格式不符合"
		c.JSON(code, data)
		return
	}

	// 已存在文件直接返回url
	if uploadType.CheckExist(src) {
		data["url"] = uploadType.GetFullUrl(fName)
		c.JSON(code, data)
		return
	}
	err = uploadType.MakePath(savePath)
	if err != nil {
		data["message"] = "文件目录创建失败, " + err.Error()
		c.JSON(code, data)
		logging.Error(err.Error())
		return
	}

	if err = c.SaveUploadedFile(fHeader, src); err != nil {
		data["message"] = "文件保存失败, " + err.Error()
		c.JSON(code, data)
		logging.Error(err.Error())
		return
	}

	data["url"] = uploadType.GetName(fName)
	c.JSON(code, data)
}

func QiNiuToken(c *gin.Context) {
	var respData = make(map[string]interface{})
	respData["code"] = 200

	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      callbackUrl,
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: "application/json",
	}
	mac := qbox.NewMac(accessKey, secretKey)
	uploadToken := putPolicy.UploadToken(mac)
	respData["token"] = uploadToken
	c.JSON(http.StatusOK, respData)
	return
}

func QiNiuUpload(c *gin.Context) {
	var respData = make(map[string]interface{})
	respData["code"] = 200
	_, fHeader, err := c.Request.FormFile("file")
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = err.Error()
		logging.Error(err.Error())
		return
	}

	uploadFileKey := c.PostForm("type") + "/" + fHeader.Filename
	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		CallbackURL:      callbackUrl,
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
		CallbackBodyType: "application/json",
	}

	mac := qbox.NewMac(accessKey, secretKey)
	uploadToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": fHeader.Filename,
		},
	}

	file, err := fHeader.Open()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = err.Error()
		logging.Error(err.Error())
		return
	}
	dataLen := int64(len(data))
	err = formUploader.Put(context.Background(), &ret, uploadToken, uploadFileKey, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		respData["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, respData)
		logging.Error(err.Error())
		return
	}
	respData["url"] = url + "/" + ret.Key
	c.JSON(http.StatusOK, respData)
	return
}

func QiNiuCallBack(c *gin.Context) {
	
}
