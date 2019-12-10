package router

import (
    "github.com/gin-gonic/gin"
    "laojgo/app/controller"
)

func Setup() *gin.Engine {
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
        v1.GET("/ping", controller.Ping)
    }


    tag := r.Group("/api/tag")
    {
        t := &controller.TagController{}
        tag.GET("/", t.Get)
        tag.POST("/", t.Add)
        tag.PUT("/:id", t.Edit)
        tag.DELETE("/:id", t.Delete)
    }

    return r
}
