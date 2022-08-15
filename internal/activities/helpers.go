package activities

import (
	"net/http"
	v1 "file-uploader/webserver/http/responses/v1"

	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, v1.MainStructure{
		Meta: v1.Meta{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Success: false,
		},
		Data: nil,
	})
}

func RespondValidationError(c *gin.Context, errs map[string]string) {
	c.JSON(http.StatusUnprocessableEntity, v1.MainStructure{
		Meta: v1.Meta{
			Code: http.StatusUnprocessableEntity,
			Message: "validation error",
			Success: false,
		},
		Data: errs,
	})
}
