package routers

import (
	"github.com/gin-gonic/gin"
	"laojgo/app/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// r.Static("/public", "./public") // 静态文件服务
	// r.LoadHTMLGlob("views/**/*") // 载入html模板目录

	// web路由
	// r.GET("/", Controllers.Home)
	// r.GET("/about", Controllers.About)
	// r.GET("/post/index", Controllers.Post)

	// 简单的路由组: v1
	v1 := r.Group("/api")
	{
		v1.GET("/ping", controllers.Ping)
        v1.GET("/tag", (&controllers.TagController{}).Get)
	}

	return r
}
