package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/pkg/upload"
)

func Upload(c *gin.Context)  {
	file, fHeader, err := c.Request.FormFile(c.PostForm("name"))

	var code = http.StatusOK
	data := make(map[string]interface{})
	if err != nil || fHeader == nil{
		code = http.StatusBadRequest
		data["message"] = "文件上传错误:"+ err.Error()
	}

	fName := upload.GetName(fHeader.Filename)
	savePath := upload.GetPath()

	src := savePath + fName

	if !upload.CheckExt(fName) || !upload.CheckSize(file) {
		code = http.StatusForbidden
		data["message"] = "文件格式不符合"
	}

	// 已存在文件直接返回url
	if upload.CheckExist(src) {
		code = http.StatusOK
		data["url"] = upload.GetFullUrl(fName)
		c.JSON(code,data)
		return
	}
	err = upload.MakePath(savePath)
	if err != nil {
		code = http.StatusInternalServerError
		data["message"] = "文件目录创建失败"
	}

	if err = c.SaveUploadedFile(fHeader, src); err != nil{
		code = http.StatusInternalServerError
		data["message"] = "文件保存失败"
	}

	data["url"] = upload.GetName(fName)
	c.JSON(code,data)
}
