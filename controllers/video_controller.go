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

	typeOfVideo := checkType(c.Query("type"))
	if typeOfVideo == -1{
		respData["code"] = http.StatusBadRequest
		respData["message"] = "请求参数错误"
		c.JSON(http.StatusBadRequest, respData)
		return
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
	typeOfVideo := checkType(c.PostForm("type"))
	if typeOfVideo == -1{
		respData["code"] = http.StatusBadRequest
		respData["message"] = "请求参数错误"
		c.JSON(http.StatusBadRequest, respData)
		return
	}

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

func VideoView(c *gin.Context) {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	// 条件
	maps := make(map[string]interface{})
	id := com.StrTo(c.Query("id")).MustInt()

	exist := models.GetVideoById(id)
	if !exist {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "数据不存在"
		c.JSON(http.StatusBadRequest, respData)
		return
	}

	maps["id"] = id
	respData["info"] = models.GetVideoView(maps)
	c.JSON(http.StatusOK, respData)
}

func VideoUpdate(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	name := c.PostForm("name")
	href := c.PostForm("href")
	typeOfVideo := checkType(c.PostForm("type"))
	if typeOfVideo == -1{
		respData["code"] = http.StatusBadRequest
		respData["message"] = "请求参数错误"
		c.JSON(http.StatusBadRequest, respData)
		return
	}

	id := com.StrTo(c.PostForm("id")).MustInt()
	exist := models.GetVideoById(id)
	if !exist {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "数据不存在"
		c.JSON(http.StatusBadRequest, respData)
		return
	}

	video := models.Video{
		Name:name,
		Href:href,
		Type:typeOfVideo,
	}

	err := models.PutVideoUpdate(id, video)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "修改数据失败," + err.Error()
		c.JSON(http.StatusBadRequest, respData)
		return
	}
	respData["message"] = "修改成功"
	c.JSON(http.StatusOK, respData)
}

func VideoDelete(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	id := com.StrTo(c.Query("id")).MustInt()
	exist := models.GetVideoById(id)
	if !exist {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "数据不存在"
		c.JSON(http.StatusBadRequest, respData)
		return
	}
	respData["message"] = "删除成功"
	c.JSON(http.StatusOK, respData)
}

func checkType(typeOfVideoName string) (typeOfVideo int){
	if typeOfVideoName == "classic"{
		typeOfVideo = 0
	} else if typeOfVideoName == "anime"{
		typeOfVideo = 1
	} else {
		typeOfVideo = -1
	}
	return typeOfVideo
}
