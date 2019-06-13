package app

import (
	"gintest/util/rs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	item := c.Query("content")
	go rs.Produce(item)
	c.JSON(http.StatusOK, "success")

}
