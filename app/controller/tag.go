package controller

import (
    "strconv"
    "net/http"

    "github.com/gin-gonic/gin"
    
    "laojgo/app/model"
    "laojgo/pkg/e"
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

    tag := &model.Tag{}
    data["total"] = tag.GetTotal(where)
    data["list"] = tag.Get(offset, limit, where)

    c.JSON(http.StatusOK, gin.H{
        "code" : e.SUCCESS,
        "msg" : e.Msg(e.SUCCESS),
        "data" : data,
    })
}

//新增文章标签
func (t *TagController) Add(c *gin.Context) {
    var tag model.Tag
    code := e.SUCCESS
    msg := e.Msg(code)

    defer c.JSON(http.StatusOK, gin.H{
        "code" : &code,
        "msg" : &msg,
    })

    if err := c.ShouldBind(&tag); err != nil {
        code = e.INVALID_PARAMS
        msg = err.Error()

        return 
    }

    if tag.ExistByName(tag.Name) {
        code = e.RESOURCE_NOT_EXIST
        msg = e.Msg(code)

        return
    }

    tag.Add()
}

//修改文章标签
func (t *TagController) Edit(c *gin.Context) {
}

//删除文章标签
func (t *TagController) Delete(c *gin.Context) {
}