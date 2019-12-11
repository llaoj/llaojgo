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

    var tag model.Tag
    data["total"] = tag.GetTotal(where)
    data["list"] = tag.Get(offset, limit, where)

    c.JSON(http.StatusOK, gin.H{
        "code" : e.SUCCESS,
        "msg" : e.Msg(e.SUCCESS),
        "data" : data,
    })
}

type createForm struct {
    Name string `form:"name" binding:"required","max=10"`
}

//新增文章标签
func (t *TagController) Create(c *gin.Context) {
    code := e.SUCCESS
    msg := e.Msg(code)

    defer c.JSON(http.StatusOK, gin.H{
        "code" : &code,
        "msg" : &msg,
    })

    var form createForm
    if err := c.ShouldBind(&form); err != nil {
        code = e.INVALID_PARAMS
        msg = err.Error()

        return 
    }

    var tag model.Tag
    if tag.ExistByName(form.Name) {
        code = e.RESOURCE_EXIST
        msg = e.Msg(code)

        return
    }

    tag.Name = form.Name
    tag.Create()
}

//删除文章标签
func (t *TagController) Delete(c *gin.Context) {
}

type updateForm struct {
    ID string `form:"id" binding:"required"`
    Name string `form:"name" binding:"required","max=10"`
}

//修改文章标签
func (t *TagController) Update(c *gin.Context) {
    code := e.SUCCESS
    msg := e.Msg(code)
    defer c.JSON(http.StatusOK, gin.H{
        "code" : &code,
        "msg" : &msg,
    })

    var form updateForm
    if err := c.ShouldBind(&form); err != nil {
        code = e.INVALID_PARAMS
        msg = err.Error()

        return 
    }

    var tag model.Tag
    tagId, _ := strconv.Atoi(form.ID)
    if !tag.ExistByID(tagId) {
        code = e.RESOURCE_NOT_EXIST
        msg = e.Msg(code)

        return
    }

    tag.Update(form)
}