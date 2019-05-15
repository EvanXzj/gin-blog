// Package util contains some help function
package util

import (
	"github.com/EvanXzj/gin-blog/pkg/setting"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// GetOffset get page parameters
func GetOffset(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
