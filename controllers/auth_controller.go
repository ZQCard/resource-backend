package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resource-backend/models"
	"resource-backend/utils"
)

func Login(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	data := make(map[string]interface{})

	isExist, _ := models.CheckAuth(username, password)
	if !isExist {
		data["message"] = "用户不存在"
		c.JSON(http.StatusBadRequest, data)
		return
	}

	token, err := utils.GenerateToken(username, password)
	if err != nil {
		data["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, data)
		return
	}
	data["token"] = token
	c.JSON(http.StatusOK, data)
}

func UserInfo(c *gin.Context)  {
	clamis, err := utils.ParseToken(c.Query("token"))
	if err != nil{
		log.Fatal(err.Error())
	}

	data := make(map[string]interface{})
	isExist, _ := models.CheckAuth(clamis.Username, clamis.Password)
	if !isExist {
		data["message"] = "用户不存在"
		c.JSON(http.StatusBadRequest, data)
		return
	}

	user := models.GetUserInfo(clamis.Username, clamis.Password)
	data["user"] =  user
	c.JSON(http.StatusOK, data)
}

