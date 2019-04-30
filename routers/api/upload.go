package api

import (
	"net/http"

	"github.com/EvanXzj/gin-blog/pkg/cerror"
	"github.com/EvanXzj/gin-blog/pkg/logging"
	"github.com/EvanXzj/gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	code := cerror.SUCCESS
	data := make(map[string]string)

	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = cerror.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  cerror.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = cerror.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(src) || !upload.CheckImageSize(file) {
			code = cerror.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = cerror.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = cerror.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  cerror.GetMsg(code),
		"data": data,
	})
}
