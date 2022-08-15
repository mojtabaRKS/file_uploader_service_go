package v1

import (
	helpers "file-uploader/internal/activities"
	"file-uploader/internal/activities/logger"
	workflows "file-uploader/internal/workflows"
	v1_response "file-uploader/webserver/http/responses/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadController(c *gin.Context) {

	model, err := workflows.Upload(c)

	if err != nil {
		logger.Logger().Error(err)
		helpers.RespondError(c, err)
		return
	}

	response := v1_response.MainStructure{
		Meta: v1_response.Meta{
			Code:    http.StatusOK,
			Message: "file_uploaded successfully",
			Success: true,
		},
		Data: v1_response.UploadResponse{
			ID:  model.ID,
			Url: model.Path,
		},
	}

	c.JSON(http.StatusOK, response)
}
