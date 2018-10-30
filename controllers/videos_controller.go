package controllers

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"resource-backend/models"
	"resource-backend/pkg/config"
)

func Videos(c *gin.Context) {
	// 错误信息
	var err error
	// 条件
	maps := make(map[string]interface{})
	// 返回数据
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	typeOfVideo := c.Query("type")
	if typeOfVideo == "" {
		respData["code"] = http.StatusBadRequest
		respData["message"] = "请求参数错误"
		c.JSON(http.StatusBadRequest, respData)
		return
	} else {
		if typeOfVideo == "classic"{
			maps["type"] = 0
		} else if typeOfVideo == "anime"{
			maps["type"] = 1
		} else {
			respData["code"] = http.StatusBadRequest
			respData["message"] = "请求参数错误"
			c.JSON(http.StatusBadRequest, respData)
			return
		}
	}

	// 分页参数
	PageSize := com.StrTo(c.DefaultQuery("pageSize", config.AppSetting.PageSize)).MustInt()
	PageNum := com.StrTo(c.DefaultQuery("page", config.AppSetting.PageNum)).MustInt()

	// 获取数据列表
	respData["list"], err = models.GetVideosList(PageNum ,PageSize, maps)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = err
		c.JSON(http.StatusBadRequest, respData)
		return
	}

	// 获取数据总记录数
	respData["totalCount"] = models.GetVideosTotalCount(maps)

	c.JSON(http.StatusOK, respData)
}

func VideoAdd(c *gin.Context)  {

	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	name := c.PostForm("name")
	href := c.PostForm("href")
	typeOfVideoName := c.PostForm("type")

	var typeOfVideo int
	if typeOfVideoName == "classic"{
		typeOfVideo = 0
	} else if typeOfVideoName == "anime"{
		typeOfVideo = 1
	} else {
		respData["code"] = http.StatusBadRequest
		respData["message"] = "请求参数错误"
		c.JSON(http.StatusBadRequest, respData)
		return
	}

	// 数据验证

	err := models.AddVideo(name, href, typeOfVideo)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "添加数据失败," + err.Error()
		c.JSON(http.StatusBadRequest, respData)
		return
	}
	respData["message"] = "添加成功"
	c.JSON(http.StatusOK, respData)
}
