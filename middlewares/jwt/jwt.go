package jwt

import (
	"net/http"
	"time"

	"github.com/EvanXzj/gin-blog/pkg/cerror"
	"github.com/EvanXzj/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
)

// JWT token verify middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = cerror.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = cerror.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = cerror.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = cerror.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != cerror.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  cerror.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
