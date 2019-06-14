package cmd

import (
	"gintest/app"
	"gintest/util/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// r := gin.New()
	r.Use(cors.Default()) //跨域

	{
		r.GET("/app/test", app.Index)
		r.GET("/app/test/del", app.DeletePerson)
		r.GET("/app/test/upd", app.UpdatePerson)

		//post
		r.POST("/app/test/add", app.AddPerson)

		r.GET("/app/interest", app.FindInterest)
		r.POST("/app/notice", app.Test)
	}
	//ws
	{
		r.GET("/ws", ws.ServeWs)
	}

	return r
}
