package routers

import (
	"github.com/EvanXzj/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

// InitRouter initiate routers
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello test",
		})
	})

	return r
}
