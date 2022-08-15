package workflows

import (
	"file-uploader/database/migrations"
	"file-uploader/internal/activities/file_uploader"
	"file-uploader/internal/activities/logger"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) (*migrations.File , error) {
	file, err := c.FormFile("file")

	if err != nil {
		logger.Logger().Error(err)
		return nil, err
	}

	err = file_uploader.Upload(file, c)

	if err != nil {
		logger.Logger().Error(err)
		return nil, err
	}

	fileModel, err := file_uploader.Save(file.Filename, int(file.Size))

	if err != nil {
		logger.Logger().Error(err)
		return nil, err
	}

	return &fileModel, nil
}