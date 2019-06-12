package cmd

import (
	"gintest/app"
	"gintest/util/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default()) //跨域

	{
		router.GET("/app/test", app.Index)
		router.GET("/app/test/del", app.DeletePerson)
		router.GET("/app/test/upd", app.UpdatePerson)

		//post
		router.POST("/app/test/add", app.AddPerson)

		router.GET("/app/interest", app.FindInterest)
		router.GET("/app/notice", app.Test)
	}
	//ws
	{
		router.GET("/ws", ws.ServeWs)
	}

	return router
}
