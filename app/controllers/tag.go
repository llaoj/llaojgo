package controllers

import (
    "github.com/gin-gonic/gin"
    "strconv"
    "laojgo/app/models"
)

type TagController struct {}

func (t *TagController) Get(c *gin.Context) {
    
    where := make(map[string]interface{})
    data := make(map[string]interface{})

    where["name"] = c.DefaultQuery("name", "")
    where["state"], _ = strconv.Atoi(c.DefaultQuery("state", "-1"))

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
    offset := (page - 1) * limit

    tag := models.Tag{}
    data["total"] = tag.GetTotal(where)
    data["list"] = tag.Get(offset, limit, where)

    c.JSON(200, gin.H{
        "code" : 0,
        "msg" : "ok",
        "data" : data,
    })
}