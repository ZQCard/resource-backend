package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/models"
	"resource-backend/pkg/logging"
	"resource-backend/utils"
)

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
		respData["userRoles"] = models.FindRoleByUserId(user.ID)

		// 找出每个角色拥有的路由和未拥有的路由
		roleRoute := make(map[string]map[string][]string)
		for _, v := range roles{
			// 临时存放
			temp := make(map[string][]string)
			temp["yes"] = models.FindRoutesByRole(v)
			temp["no"] = trimRoutes(routers, temp["yes"])
			roleRoute[v] = temp
		}
		respData["roleRoutes"] = roleRoute

		c.JSON(http.StatusOK, respData)
		return
	}
}

func trimRoutes(all []string, yes []string) (no []string) {
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
