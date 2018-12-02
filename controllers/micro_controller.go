package controllers

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/models"
	"resource-backend/pkg/config"
	"resource-backend/pkg/logging"
)

func MicroVideoList(c *gin.Context) {
	// 错误信息
	var err error
	// 返回数据
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	// 分页参数
	PageSize := com.StrTo(c.DefaultQuery("pageSize", config.AppSetting.PageSize)).MustInt()
	PageNum := com.StrTo(c.DefaultQuery("page", config.AppSetting.PageNum)).MustInt()

	// 获取数据列表
	respData["list"], respData["totalCount"], err = models.MicroVideoList(PageNum ,PageSize)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = err.Error()
		c.JSON(http.StatusBadRequest, respData)
		logging.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, respData)
	return
}

func MicroVideoAdd(c *gin.Context)  {

	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	url := c.PostForm("url")

	// 检测视频是否已经存在
	exist := models.FindMicroVideoByUrl(url)
	if exist == true {
		respData["code"] = http.StatusExpectationFailed
		respData["message"] = "视频已存在"
		c.JSON(http.StatusExpectationFailed, respData)
		return
	}
	microVideo := models.MicroVideo{
		Url:url,
	}

	err := models.AddMicroVideo(&microVideo)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "添加数据失败," + err.Error()
		c.JSON(http.StatusBadRequest, respData)
		logging.Error(err.Error())
		return
	}
	respData["message"] = "添加成功"
	c.JSON(http.StatusOK, respData)
}

func MicroVideoView(c *gin.Context)  {
	resp := make(map[string]interface{})
	resp["code"] = 200
	url := c.Query("url")
	if url == ""{
		resp["code"] = http.StatusBadRequest
		resp["message"] = "请求参数错误"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	err := models.ViewMicroVideo(url)
	if err != nil {
		resp["code"] = http.StatusInternalServerError
		resp["message"] = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp["url"] = url
	c.JSON(http.StatusOK, resp)
	return
}

func MicroVideoDelete(c *gin.Context)  {
	resp := make(map[string]interface{})
	resp["code"] = http.StatusOK

	url := com.StrTo(c.Query("url")).String()
	err := models.DeleteMicroVideo(url)
	if err != nil {
		resp["code"] = http.StatusInternalServerError
		resp["message"] = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp["message"] = "删除成功"
	c.JSON(http.StatusOK, resp)
	return
}
