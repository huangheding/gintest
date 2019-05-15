package app

import (
	"gin_test/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	p := model.Person{}
	result, _ := p.FindPerson()

	c.JSON(http.StatusOK, result)
}
