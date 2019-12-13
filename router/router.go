package router

import (
    "github.com/gin-gonic/gin"

    "laojgo/config"
    "laojgo/app/controller"
    "laojgo/app/middleware"
)

func Setup() *gin.Engine {
    gin.SetMode(config.Server.RunMode)

    // 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
    gin.DisableConsoleColor()

    // 记录到文件。
    // f, _ := os.Create("gin.log")
    // gin.DefaultWriter = io.MultiWriter(f)

    // 如果需要同时将日志写入文件和控制台，请使用以下代码。
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)


    r := gin.Default()

    // r.Static("/public", "./public") // 静态文件服务
    // r.LoadHTMLGlob("views/**/*") // 载入html模板目录

    // 接口鉴权
    auth := r.Group("/auth")
    {
        a := &controller.AuthController{}
        auth.POST("/", a.GetToken)
    }

    v1 := r.Group("/v1")
    v1.Use(middleware.Jwt())
    v1.GET("/ping", controller.Ping) //example

    tag := v1.Group("/tag")
    {
        t := &controller.TagController{}
        tag.GET("/", t.Get)
        tag.POST("/", t.Create)
        tag.PUT("/:id", t.Update)
        tag.DELETE("/:id", t.Delete)
    }

    return r
}
