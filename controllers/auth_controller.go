package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/models"
	"resource-backend/pkg/logging"
	"resource-backend/utils"
	"strings"
)

// 获取权限数据集合
func Auth(router *gin.Engine) func(c *gin.Context){
	return func(c *gin.Context) {
		respData := make(map[string]interface{})
		respData["code"] = http.StatusOK
		routers := []string{}
		for _, v := range router.Routes() {
			routers = append(routers, v.Method+":"+v.Path)
		}
		// 路由列表
		respData["routes"] = routers
		// 角色列表
		roles := models.RoleList()
		respData["roles"] = roles
		// 当前用户拥有的角色
		clamis, err := utils.ParseToken(c.Query("token"))
		if err != nil{
			respData["message"] = err.Error()
			c.JSON(http.StatusBadRequest, respData)
			logging.Error(err)
			return
		}

		maps := map[string]interface{}{
			"username":clamis.Username,
			"password":utils.EncodeMD5(clamis.Password),
		}
		user, err := models.GetUserByMaps(maps)
		if err != nil {
			respData["message"] = "用户不存在"
			c.JSON(http.StatusBadRequest, respData)
			return
		}
		userRole := make(map[string][]string)
		userRole["has"] = models.FindRoleByUserId(user.ID)
		userRole["no"] = filterDiff(roles, userRole["has"])
		respData["userRoles"] = userRole

		// 找出每个角色拥有的路由和未拥有的路由
		roleRoute := make(map[string]map[string][]string)
		for _, v := range roles{
			// 临时存放
			temp := make(map[string][]string)
			temp["yes"] = models.FindRoutesByRole(v)
			temp["no"] = filterDiff(routers, temp["yes"])
			roleRoute[v] = temp
		}
		respData["roleRoutes"] = roleRoute

		c.JSON(http.StatusOK, respData)
		return
	}
}

// 筛选出不在yes数组中all的元素素组
func filterDiff(all []string, yes []string) (no []string) {
	if yes == nil {
		return all
	}
	yesMap := make(map[string]int)
	for  k, v := range yes{
		yesMap[v] = k
	}
	for _, v := range all{
		_,ok := yesMap[v]
		if !ok {
			no = append(no, v)
		}
	}
	return
}

func Assign(c *gin.Context)  {
	resp := make(map[string]interface{})
	claims, err := utils.ParseToken(c.Query("token"))
	if err != nil {
		if err != nil{
			logging.Error(err)
			return
		}
	}

	maps := map[string]interface{}{
		"username":claims.Username,
		"password":utils.EncodeMD5(claims.Password),
	}
	user, err := models.GetUserByMaps(maps)

	if err != nil {
		resp["message"] = "用户不存在"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	roleName := c.PostForm("role")
	if !models.CheckRoleExist(roleName){
		resp["message"] = "角色不存在,请先创建角色"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	models.FindOrCreateAssignment(user.ID, roleName)
	resp["message"] = "设置成功"
	c.JSON(http.StatusOK, resp)
	return
}

func Allocate(c *gin.Context)  {
	resp := make(map[string]interface{})
	role := c.PostForm("role")
	route := c.PostForm("route")
	temp := strings.Split(route, ":")
	routeName := temp[0]
	method := temp[1]
	err := models.AddRoleRoute(role, routeName, method)
	if err != nil {
		resp["message"] = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		logging.Error(err.Error())
		return
	}
	resp["message"] = "分配成功"
	c.JSON(http.StatusOK, resp)
	return
}