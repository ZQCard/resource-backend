package controllers

import (
	"github.com/Unknwon/com"
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
		claims, err := utils.ParseToken(c.Query("token"))
		if err != nil{
			respData["message"] = err.Error()
			c.JSON(http.StatusBadRequest, respData)
			logging.Error(err)
			return
		}

		maps := map[string]interface{}{
			"username":claims.Username,
			"password":utils.EncodeMD5(claims.Password),
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

	maps := map[string]interface{}{
		"id":com.StrTo(c.Query("id")).MustInt(),
	}
	user, err := models.GetUserByMaps(maps)
	if err != nil {
		resp["message"] = "用户不存在"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// 要授权的角色数组
	roleNames := strings.Split(c.PostForm("roles"), ",")
	for _,role := range roleNames{
		if !models.CheckRoleExist(role){
			resp["message"] = "角色 "+role+" 不存在,请先创建角色"
			c.JSON(http.StatusBadRequest, resp)
			return
		}
	}
	// 已经拥有的角色
	rolesHas := models.FindRoleByUserId(user.ID)
	rolesHasNot := filterDiff(roleNames,rolesHas)
	for _,role := range rolesHasNot{
		models.AssignRoles(user.ID, role)
	}
	resp["message"] = "设置成功"
	c.JSON(http.StatusOK, resp)
	return
}

func Allocate(c *gin.Context)  {
	resp := make(map[string]interface{})
	role := c.PostForm("role")
	route := c.PostForm("route")
	temp := strings.Split(route, ":")
	method := temp[0]
	routeName := temp[1]
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

func Assignment(c *gin.Context)  {
	resp := make(map[string]interface{})
	resp["code"] = http.StatusOK

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

	id := com.StrTo(c.Query("id")).MustInt()
	if id == 0 {
		resp["code"] = http.StatusBadRequest
		resp["message"] = "参数错误!"
		c.JSON(http.StatusInternalServerError, resp)
		logging.Error("授权参数错误"+user.Username)
		return
	}
	// 角色列表
	roles := models.RoleList()
	resp["roles"] = roles

	// 查找所有的和拥有的角色
	userRole := make(map[string][]string)
	userRole["has"] = models.FindRoleByUserId(user.ID)
	userRole["no"] = filterDiff(roles, userRole["has"])
	resp["userRoles"] = userRole
	c.JSON(http.StatusOK, resp)
	return
}