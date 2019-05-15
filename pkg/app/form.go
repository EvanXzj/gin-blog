package app

import (
	"net/http"

	"github.com/EvanXzj/gin-blog/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	ok, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}

	if !ok {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
