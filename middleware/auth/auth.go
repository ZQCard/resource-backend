package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/models"
	"resource-backend/utils"
	"time"
)

func AUTH() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK
		// 查询当前用户所拥有的路由
		claims, err := utils.ParseToken(c.Query("token"))
		if err != nil {
			code = http.StatusBadRequest
		}else if time.Now().Unix() > claims.ExpiresAt {
			code = http.StatusRequestTimeout
		}
		maps := map[string]interface{}{
			"username":claims.Username,
			"password":utils.EncodeMD5(claims.Password),
		}
		user, err := models.GetUserByMaps(maps)
		routes := models.GetRoutesByUserId(user.ID)
		flag := false
		for  _, v := range routes{
			if v == {

			}
		}
		c.Next()
	}
}
