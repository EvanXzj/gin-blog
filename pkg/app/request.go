package app

import (
	"github.com/EvanXzj/gin-blog/pkg/logging"
	"github.com/astaxie/beego/validation"
)

// MarkErrors logs error log
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
