package controllers

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/models"
	"resource-backend/pkg/config"
	"resource-backend/pkg/logging"
	"strings"
	"time"
)

func PersonalBillSummary(c *gin.Context)  {
	// 返回数据
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK

	// 支出
	expandMaps := make(map[string]interface{})
	expandMaps["type"] = models.PersonalBillExpand
	// 总消费
	respData["expand_all"] = models.PersonalBillSummaryByCategory(expandMaps)

	// 年度消费
	expandMaps["year"] = time.Now().Year()
	respData["expand_year"] = models.PersonalBillSummaryByCategory(expandMaps)
	// 本月消费
	expandMaps["month"] = strings.TrimPrefix(time.Now().Format("01"), "0")
	respData["expand_month"] = models.PersonalBillSummaryByCategory(expandMaps)

	// 收入
	incomeMaps := make(map[string]interface{})
	incomeMaps["type"] = models.PersonalBillIncome
	// 总收入
	respData["income_all"] = models.PersonalBillSummaryByCategory(incomeMaps)
	// 年度收入
	incomeMaps["year"] = time.Now().Year()
	respData["income_year"] = models.PersonalBillSummaryByCategory(incomeMaps)
	// 本月收入
	incomeMaps["month"] = strings.TrimPrefix(time.Now().Format("01"), "0")
	respData["income_month"] = models.PersonalBillSummaryByCategory(incomeMaps)

	// 收入支出对比
	// 总数
	incomeExpandMaps := make(map[string]interface{})
	respData["expand_income_all"] = models.PersonalBillSummaryByType(incomeExpandMaps)
	// 年度
	incomeExpandMaps["year"] = time.Now().Year()
	respData["expand_income_year"] = models.PersonalBillSummaryByType(incomeExpandMaps)
	// 月数
	incomeExpandMaps["month"] = strings.TrimPrefix(time.Now().Format("01"), "0")
	respData["expand_income_month"] = models.PersonalBillSummaryByType(incomeExpandMaps)

	// 年度收入支出折线图
	year := c.Query("year")
	eachMonthMap := make(map[string]interface{})
	if year == "" {
		eachMonthMap["year"] = time.Now().Year()
	}else {
		eachMonthMap["year"] = year
	}
	respData["year_expand"], respData["year_income"] = models.PersonalBillSummaryByYear(eachMonthMap)



	// 总结
	respData["summary"] = models.PersonalBillSummary()
	c.JSON(http.StatusOK, respData)
	return
}

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

func PersonalBillView(c *gin.Context) {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	// 条件
	maps := make(map[string]interface{})
	maps["id"] = com.StrTo(c.Query("id")).MustInt()
	respData["info"] = models.PersonalBillView(maps)
	c.JSON(http.StatusOK, respData)
}

func PersonalBillAdd(c *gin.Context)  {
	respData := make(map[string]interface{})
	respData["code"] = http.StatusOK
	typeOfMoney := com.StrTo(c.PostForm("type")).MustInt()
	money,_ := com.StrTo(c.PostForm("money")).Float64()
	date := c.PostForm("date")
	category := c.PostForm("category")
	dateOfArray := strings.Split(date, "-")
	PersonalBill := &models.PersonalBill{
		Type:typeOfMoney,
		Money:money,
		Category:category,
		Year:dateOfArray[0],
		Month:dateOfArray[1],
		Day:dateOfArray[2],
	}

	err := models.PersonalBillAdd(PersonalBill)
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

	typeOfMoney := com.StrTo(c.PostForm("type")).MustInt()
	date := c.PostForm("date")
	dateOfArray := strings.Split(date, "-")
	maps := map[string]interface{}{
		"id":com.StrTo(c.PostForm("id")).MustInt(),
	}
	PersonalBill := models.PersonalBillView(maps)
	PersonalBill.Type = typeOfMoney
	PersonalBill.Money,_ = com.StrTo(c.PostForm("money")).Float64()
	PersonalBill.Category = c.PostForm("category")
	PersonalBill.Year = dateOfArray[0]
	PersonalBill.Month = dateOfArray[1]
	PersonalBill.Day = dateOfArray[2]
	err := models.PersonalBillUpdate(&PersonalBill)
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
