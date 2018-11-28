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

func PersonalBillList(c *gin.Context) {
	// 错误信息
	var err error
	// 返回数据
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	// 分页参数
	PageSize := com.StrTo(c.DefaultQuery("pageSize", config.AppSetting.PageSize)).MustInt()
	PageNum := com.StrTo(c.DefaultQuery("page", config.AppSetting.PageNum)).MustInt()

	// 获取数据列表
	respData["list"], respData["totalCount"], err = models.PersonalBillList(PageNum ,PageSize)
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

func PersonalBillAdd(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	PersonalBill := models.PersonalBill{

	}

	err := models.PersonalBillAdd(&PersonalBill)
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

func PersonalBillUpdate(c *gin.Context)  {
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
	user.Password = utils.EncodeMD5(c.PostForm("avatar"))

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

func PersonalBillDelete(c *gin.Context)  {
	resp := make(map[string]interface{})
	resp["code"] = http.StatusOK

	id := com.StrTo(c.Query("id")).MustInt()
	err := models.PersonalBillDelete(id)
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
