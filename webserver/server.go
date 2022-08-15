package webserver

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RunServer(routesHandler func(r *gin.Engine) (*gin.Engine), port string) {
	gin.SetMode(os.Getenv("APP_ENV"))
	apiRouter := gin.Default()
	logrus.Info("Starting HTTP server")

	err := apiRouter.Run(":" + port)
	if err != nil {
		panic(err)
	}
}