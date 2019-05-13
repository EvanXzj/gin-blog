package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/EvanXzj/gin-blog/pkg/e"
	"github.com/EvanXzj/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
)

// JWT token verify middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			// this is mandatory
			c.Abort()
			return
		}

		c.Next()
	}
}
