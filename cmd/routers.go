package cmd

import (
	"gintest/app"
	"gintest/util/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func middleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("I am before next")
// 		// c.Header("Access-Control-Allow-Origin", "test")
// 		// c.Set("name", "test")
// 		/*
// 		   c.Next()后就执行真实的路由函数，路由函数执行完成之后继续执行后续的代码
// 		*/
// 		c.Next()
// 		fmt.Println("I am after next")
// 	}
// }

func InitRouter() *gin.Engine {
	r := gin.Default()

	// r.Use(middleware())
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
