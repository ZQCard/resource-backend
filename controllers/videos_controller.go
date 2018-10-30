package controllers

import (
	"gin-crud/models"
	"gin-crud/pkg/config"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
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

	// 数据处理
	c.JSON(http.StatusOK, respData)
}
