package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-backend/controllers"
	"resource-backend/middleware/auth"
	"resource-backend/middleware/cors"
	"resource-backend/middleware/jwt"
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

	api := router.Group("/")
	// JWT
	api.Use(jwt.JWT())
	// 权限控制
	api.Use(auth.AUTH())
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
		// 权限控制
		api.POST("/assign", controllers.Assign)
		api.POST("/allocate", controllers.Allocate)
		api.GET("/auth", controllers.Auth(router))

		// 文件上传
		api.POST("/upload", controllers.Upload)
	}
	return router
}
