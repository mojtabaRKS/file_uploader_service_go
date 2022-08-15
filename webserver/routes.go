package webserver

import (
	"github.com/gin-gonic/gin"
	v1_controller "file-uploader/webserver/http/controllers/v1"
)

func Routes(r *gin.Engine) *gin.Engine {
	r.POST("upload", v1_controller.UploadController)

	return r
}