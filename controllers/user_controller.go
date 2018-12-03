package controllers

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/models"
	"resource-backend/pkg/config"
	"resource-backend/pkg/logging"
	"resource-backend/utils"
)

func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	data := make(map[string]interface{})

	maps := map[string]interface{}{
		"username":username,
		"password":utils.EncodeMD5(password),
	}
	_, err := models.GetUserByMaps(maps)
	if err != nil {
		data["message"] = "用户名或者密码错误"
		c.JSON(http.StatusBadRequest, data)
		return
	}

	token, err := utils.GenerateToken(username, password)
	if err != nil {
		data["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, data)
		logging.Error(err)
		return
	}
	data["token"] = token
	c.JSON(http.StatusOK, data)
}

func UserList(c *gin.Context)  {
	// 错误信息
	var err error

	// 返回数据
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	// 分页参数
	PageSize := com.StrTo(c.DefaultQuery("pageSize", config.AppSetting.PageSize)).MustInt()
	PageNum := com.StrTo(c.DefaultQuery("page", config.AppSetting.PageNum)).MustInt()

	// 获取数据列表
	respData["list"], respData["totalCount"], err = models.UserList(PageNum ,PageSize)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = err.Error()
		c.JSON(http.StatusBadRequest, respData)
		logging.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, respData)
}

func UserInfo(c *gin.Context)  {
	claims, err := utils.ParseToken(c.Query("token"))
	if err != nil{
		logging.Error(err)
		return
	}

	data := make(map[string]interface{})

	maps := map[string]interface{}{
		"username":claims.Username,
		"password":utils.EncodeMD5(claims.Password),
	}
	if c.Query("id") != ""{
		maps = map[string]interface{}{
			"id":com.StrTo(c.Query("id")).MustInt(),
		}
	}
	user, err := models.GetUserByMaps(maps)

	if err != nil {
		data["message"] = "用户不存在"
		c.JSON(http.StatusBadRequest, data)
		return
	}

	data["user"] =  user
	c.JSON(http.StatusOK, data)
}

func UserAdd(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	username := c.PostForm("username")
	maps := map[string]interface{}{
		"username":username,
	}
	_, err := models.GetUserByMaps(maps)
	if err == nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "用户名已经存在"
		c.JSON(http.StatusBadRequest, respData)
		return
	}

	user := &models.User{
		Username:c.PostForm("username"),
		Password:utils.EncodeMD5(c.PostForm("password")),
		Avatar:c.PostForm("avatar"),
	}

	err = models.AddUser(user)
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

func UserUpdate(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	maps := map[string]interface{}{
		"id":com.StrTo(c.PostForm("id")).MustInt(),
	}
	user, err := models.GetUserByMaps(maps)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = err.Error()
		c.JSON(http.StatusBadRequest, respData)
		return
	}
	user.Avatar = c.PostForm("avatar")
	password := c.PostForm("password")
	if password != ""{
		user.Password = utils.EncodeMD5(password)
	}
	err = models.UserUpdate(&user)
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

func UserDelete(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	id := com.StrTo(c.Query("id")).MustInt()
	err := models.UserDelete(id)
	if err != nil {
		respData["code"] = http.StatusInternalServerError
		respData["message"] = "删除数据失败," + err.Error()
		c.JSON(http.StatusBadRequest, respData)
		logging.Error(err.Error())
		return
	}
	respData["message"] = "删除成功"
	c.JSON(http.StatusOK, respData)
}

func UserRecover(c *gin.Context) {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	id := com.StrTo(c.Query("id")).MustInt()
	models.UserRecover(id)
	respData["message"] = "恢复成功"
	c.JSON(http.StatusOK, respData)
	return
}
