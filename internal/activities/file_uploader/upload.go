package file_uploader

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func Upload(file *multipart.FileHeader, c *gin.Context) error {
	
	c.SaveUploadedFile(file, "storage/files")

	return nil
}