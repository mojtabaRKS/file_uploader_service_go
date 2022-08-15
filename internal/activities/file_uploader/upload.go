package file_uploader

import (
	"file-uploader/internal/activities/logger"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func Upload(file *multipart.FileHeader, c *gin.Context) error {
	err := c.SaveUploadedFile(file, "storage/files/" + file.Filename)

	if err != nil {
		logger.Logger().Error(err)
		return err
	}

	return nil
}