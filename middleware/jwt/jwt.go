package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/pkg/logging"
	"resource-backend/utils"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = http.StatusOK

		token := c.DefaultQuery("token", "1.1.1")

		if token == "" {
			code = http.StatusUnauthorized
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = http.StatusBadRequest
			logging.Error(err.Error())
		}else if time.Now().Unix() > claims.ExpiresAt {
			code = http.StatusRequestTimeout
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"message" : "用户验证失败",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}


