package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/controllers"
	"resource-backend/middleware/cors"
	"resource-backend/middleware/jwt"
	"resource-backend/models"
)

func InitRouter() *gin.Engine {
	// 初始化router
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	// 使用跨域中间件
	router.Use(cors.Cors())

	// 静态文件访问
	router.StaticFS("/static", http.Dir("static"))
	// 获取token
	router.POST("/login", controllers.Login)

	// 权限控制
	// 将所有路由存储到数据表中
	router.GET("/routers/refresh", func(c *gin.Context) {
		// 存储到数据库中
		var routes = []models.Routes{}

		for _, v := range router.Routes() {
			r := models.Routes{}
			r.Method = v.Method
			r.Path = v.Path
			routes = append(routes, r)
		}
		models.Refresh(routes)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/auth", controllers.Auth(router))

	// 权限控制
	api := router.Group("/")
	api.Use(jwt.JWT())
	{
		// 用户列表
		api.GET("/users", controllers.UserList)
		// 用户信息
		api.GET("/user", controllers.UserInfo)
		// 添加用户
		api.POST("/user", controllers.UserAdd)
		// 更新用户信息
		api.PUT("/user", controllers.UserUpdate)
		// 删除用户
		api.DELETE("/user", controllers.UserDelete)
		// 恢复用户
		api.PATCH("/user/recover", controllers.UserRecover)
		// 影视专区
		api.GET("/videos", controllers.Videos)
		// 添加视频
		api.POST("/video", controllers.VideoAdd)
		// 视频详情
		api.GET("/video", controllers.VideoView)
		// 更新视频
		api.PUT("/video", controllers.VideoUpdate)
		// 删除视频
		api.DELETE("/video", controllers.VideoDelete)

		// 文件上传
		api.POST("/upload", controllers.Upload)
	}
	return router
}
