package controllers

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/models"
	"resource-backend/pkg/config"
	"resource-backend/pkg/logging"
)

func PersonalDiaryList(c *gin.Context) {
	// 错误信息
	var err error
	// 返回数据
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	// 分页参数
	PageSize := com.StrTo(c.DefaultQuery("pageSize", config.AppSetting.PageSize)).MustInt()
	PageNum := com.StrTo(c.DefaultQuery("page", config.AppSetting.PageNum)).MustInt()

	// 获取数据列表
	respData["list"], respData["totalCount"], err = models.PersonalDiaryList(PageNum ,PageSize)
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

func PersonalDiaryAdd(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	PersonalDiary := &models.PersonalDiary{
		Title:c.PostForm("title"),
		Content:c.PostForm("content"),
		Secret:c.PostForm("secret"),
	}

	err := models.PersonalDiaryAdd(PersonalDiary)
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

func PersonalDiaryView(c *gin.Context) {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	// 条件
	maps := make(map[string]interface{})
	maps["id"] = com.StrTo(c.Query("id")).MustInt()
	diary := models.PersonalDiaryView(maps)
	secret := c.Query("secret")
	if secret == "" || diary.Secret != secret {
		respData["code"] = http.StatusNotFound
		respData["message"] = "密码错误"
		c.JSON(http.StatusBadRequest, respData)
		return
	}
	return
	respData["info"] = diary
	c.JSON(http.StatusOK, respData)
}

func PersonalDiaryUpdate(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	maps := make(map[string]interface{})
	maps["id"] = 3
	PersonalDiary := models.PersonalDiaryView(maps)
	PersonalDiary.Title = c.PostForm("title")
	PersonalDiary.Content = c.PostForm("content")
	PersonalDiary.Secret = c.PostForm("secret")
	err := models.PersonalDiaryUpdate(&PersonalDiary)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "修改数据失败," + err.Error()
		c.JSON(http.StatusBadRequest, respData)
		logging.Error(err.Error())
		return
	}
	respData["message"] = "修改成功"
	c.JSON(http.StatusOK, respData)
}

func PersonalDiaryDelete(c *gin.Context)  {
	resp := make(map[string]interface{})
	resp["code"] = http.StatusOK

	id := com.StrTo(c.Query("id")).MustInt()
	err := models.PersonalDiaryDelete(id)
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
