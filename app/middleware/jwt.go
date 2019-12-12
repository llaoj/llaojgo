package middleware

import (
    "time"
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
        } else {
            claims, err := util.ParseJwt(token)
            if err != nil || time.Now().Unix() > claims.ExpiresAt {
                code = e.TOKEN_INVALID
            }
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