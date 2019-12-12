package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "laojgo/pkg/util"
    "laojgo/pkg/e"

)

func Jwt() gin.HandlerFunc {
    return func(c *gin.Context) {
        code := e.SUCCESS
        token := c.Query("token")
        if token == "" {
            code = e.INVALID_PARAMS
        } else if _, err := util.ParseJwt(token); err != nil {
            code = e.AUTH_TOKEN_INVALID
        }
        if code != e.SUCCESS {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code" : code,
                "msg" : e.Msg(code),
            })
            c.Abort()
            return
        }

        c.Next()
    }
}