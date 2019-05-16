package cmd

import (
	"gintest/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default()) //跨域

	{
		router.GET("/app/test", app.Index)
	}

	return router
}
