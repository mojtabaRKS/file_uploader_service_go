package main

import (
	"fmt"
	"os"

	"file-uploader/webserver"
	"file-uploader/database/migrations"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic("load .env encountered an error")
	}
}

func main() {
	migrations.Migrate()
	
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	routeHandler := func(r *gin.Engine) *gin.Engine {
		return webserver.Routes(r)
	}

	webserver.RunServer(routeHandler, os.Getenv("APP_PORT"))
}
