package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "laojgo/app/model"
    "laojgo/pkg/util"
    "laojgo/pkg/e"
)

type AuthController struct {}


type getTokenFrom struct {
    Username string `form:"username" binding:"required"`
    Password string `form:"password" binding:"required"`
}

func (a *AuthController) GetToken(c *gin.Context) {
    code := e.SUCCESS
    msg := e.Msg(code)
    data := make(map[string]interface{})

    defer c.JSON(http.StatusOK, gin.H{
        "code" : &code,
        "msg" : &msg,
        "data" : &data,
    })

    var form getTokenFrom
    if err := c.ShouldBind(&form); err != nil {
        code = e.INVALID_PARAMS
        msg = err.Error()
        return 
    }

    where := map[string]interface{}{
        "username" : form.Username,
        "password" : form.Password,
    }

    var auth model.Auth
    if !auth.Check(where) {
        code = e.AUTH_FAIL
        msg = e.Msg(code)
        return
    }

    token, err := util.GenerateJwt(form.Username)
    if err != nil {
        code = e.AUTH_FAIL
        msg = e.Msg(code)
        return
    }

    data["token"] = token
}