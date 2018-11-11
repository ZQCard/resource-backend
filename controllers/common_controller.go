package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/pkg/logging"
	"resource-backend/pkg/upload"
)

func Upload(c *gin.Context)  {
	fType := c.PostForm("type")
	switch fType {
	case "image":
		processImage(c)
	default:
		c.JSON(http.StatusOK, gin.H{
			"message":"请求参数错误",
		})
	}
	return
}

func processImage(c *gin.Context)  {
	var code = http.StatusOK
	data := make(map[string]interface{})
	file, fHeader, err := c.Request.FormFile("image")
	if err != nil || fHeader == nil{
		data["message"] = "文件上传错误:"+ err.Error()
		c.JSON(code,data)
		logging.Error(err.Error())
		return
	}
	uploadType := &upload.Image{}

	fName := uploadType.GetName(fHeader.Filename)
	savePath := uploadType.GetPath()

	src := savePath + fName

	if !uploadType.CheckExt(fName) || !uploadType.CheckSize(file) {
		data["message"] = "文件格式不符合"
		c.JSON(code,data)
		return
	}

	// 已存在文件直接返回url
	if uploadType.CheckExist(src) {
		data["url"] = uploadType.GetFullUrl(fName)
		c.JSON(code,data)
		return
	}
	err = uploadType.MakePath(savePath)
	if err != nil {
		data["message"] = "文件目录创建失败, " + err.Error()
		c.JSON(code,data)
		logging.Error(err.Error())
		return
	}

	if err = c.SaveUploadedFile(fHeader, src); err != nil{
		data["message"] = "文件保存失败, " + err.Error()
		c.JSON(code,data)
		logging.Error(err.Error())
		return
	}

	data["url"] = uploadType.GetName(fName)
	c.JSON(code,data)
}
