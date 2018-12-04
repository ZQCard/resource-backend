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
	router.POST("/user_login", controllers.Login)

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
		// 申请成为用户
		api.POST("/user/apply", controllers.UserApply)

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

		// 微视频
		// 列表
		api.GET("/micro_videos", controllers.MicroVideoList)
		// 添加
		api.POST("/micro_video", controllers.MicroVideoAdd)
		// 观看
		api.GET("/micro_video", controllers.MicroVideoView)
		// 删除
		api.DELETE("/micro_video", controllers.MicroVideoDelete)

		// 权限控制
		// 角色授权情况
		api.GET("/assignment", controllers.Assignment)
		// 角色授予
		api.POST("/assign", controllers.Assign)
		// 角色路由分配情况
		api.GET("/auth", controllers.Auth(router))
		// 路由分配
		api.POST("/role", controllers.Allocate)
		// 删除角色
		api.DELETE("/role", controllers.RoleRemove)

		// 个人隐私
		// 照片相关
		// 列表
		api.GET("/personal/photos", controllers.PersonalPhotoList)
		// 添加
		api.POST("/personal/photo", controllers.PersonalPhotoAdd)
		// 删除
		api.DELETE("/personal/photo", controllers.PersonalPhotoDelete)

		// 账单相关
		// 金额统计图
		// api.GET("/personal/summary", controllers.PersonalPhotoList)
		// 列表
		api.GET("/personal/bills", controllers.PersonalBillList)
		// 添加
		api.POST("/personal/bill", controllers.PersonalBillAdd)
		// 详情
		api.GET("/personal/bill", controllers.PersonalBillView)
		// 编辑
		api.PUT("/personal/bill", controllers.PersonalBillUpdate)
		// 删除
		api.DELETE("/personal/bill", controllers.PersonalBillDelete)

		// 日记相关
		api.GET("/personal/diaries", controllers.PersonalDiaryList)
		// 添加
		api.POST("/personal/diary", controllers.PersonalDiaryAdd)
		// 详情
		api.GET("/personal/diary", controllers.PersonalDiaryView)
		// 编辑
		api.PUT("/personal/diary", controllers.PersonalDiaryUpdate)
		// 删除
		api.DELETE("/personal/diary", controllers.PersonalDiaryDelete)

		// 文件上传
		api.POST("/upload", controllers.Upload)

		// 七牛云文件上传 服务器
		api.POST("/qiniu-upload", controllers.QiNiuUpload)
		// 七牛云上传的token 客户端
		api.GET("/qiniu-token", controllers.QiNiuToken)
	}
	return router
}
