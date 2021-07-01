package jwt

import (
	"github.com/Grey-12/gin-blog/pkg/errorCode"
	"github.com/Grey-12/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = errorCode.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = errorCode.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = errorCode.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errorCode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != errorCode.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg": errorCode.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}