package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/pkg/config"
	"strconv"
)

func Upload(c *gin.Context)  {
	path := "./static/"
	urlPath := "/static/"
	file, _ := c.FormFile("file")
	dst := path + file.Filename
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"url" : config.AppSetting.BaseUrl + ":" + strconv.Itoa(config.ServerSettings.HTTPPort) + urlPath + file.Filename,
	})
	return
}
