package controller

import (
    "github.com/gin-gonic/gin"
    "laojgo/pkg/log"
)

func Ping(c *gin.Context) {
    log.App.Info("A group of walrus emerges from the ocean") // applog
  
    c.JSON(200, gin.H{
        "message": "pong",
    })
}




